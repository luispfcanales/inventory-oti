package mem

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/luispfcanales/inventory-oti/models"
)

type storageActor struct {
	List       map[string]*models.User
	mailbox    chan func()
	stopSignal chan struct{}
}

func NewStorage() *storageActor {
	s := &storageActor{
		List:       make(map[string]*models.User),
		mailbox:    make(chan func(), 10),
		stopSignal: make(chan struct{}, 10),
	}

	id := uuid.New()
	s.List["luispfcanales@gmail.com"] = &models.User{
		Key:      id.String(),
		Fullname: "luispfcanales@gmail.com",
		Password: "luisangel",
	}
	go s.start()
	return s
}

func (s *storageActor) start() {
	log.Println("run actor storage")
loopActor:
	for {
		select {
		case action := <-s.mailbox:
			action()
		case <-s.stopSignal:
			break loopActor
		}
	}
}

func (s *storageActor) GetUserWithCredentials(email, pwd string) (models.User, error) {
	var err error = nil
	userChan := make(chan models.User, 0)

	s.receiver(func() {
		v, ok := s.List[email]
		if !ok {
			err = errors.New("user not found in storage")
			userChan <- models.User{}
			return
		}
		if pwd != v.Password {
			err = errors.New("bad credentials")
			userChan <- models.User{}
			return
		}
		userChan <- *v
	})

	return <-userChan, err
}

// Receiver send to channel mailbox function to execute concurret
func (s *storageActor) receiver(actionChan func()) {
	s.mailbox <- actionChan
}
