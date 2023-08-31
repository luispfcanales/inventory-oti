package clouddeta

import (
	"errors"
	"log"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/luispfcanales/inventory-oti/models"
)

type detaCloud struct {
	client *deta.Deta
}

func NewCloudDetaStorage(key string) *detaCloud {
	d, err := deta.New(deta.WithProjectKey(key))
	if err != nil {
		log.Fatal("Error Storage -> detaspace: ", err)
	}

	cloud := &detaCloud{
		client: d,
	}

	return cloud
}
func (s *detaCloud) logError(text string, err error) {
	if err != nil {
		log.Fatalf("STORAGE: %s -> %v :", text, err)
	}
}

func (s *detaCloud) GetComputers() []*models.CPU {
	var result []*models.CPU

	db, err := base.New(s.client, "equipos")
	s.logError("failed init base instance", err)

	_, err = db.Fetch(&base.FetchInput{
		Dest: &result,
	})
	s.logError("failed to fetch items", err)

	return result
}

// implement StorageUserService
func (s *detaCloud) existsEmail(email string) bool {
	var result []models.User
	db, err := base.New(s.client, "usuario")
	s.logError("failed init base instance", err)

	query := base.Query{
		{
			"email": email,
		},
	}
	_, err = db.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})
	s.logError("failed to fetch items", err)
	log.Println("exist email: ", result)
	log.Println("exist email: ", len(result))

	if len(result) > 0 {
		for _, value := range result {
			log.Println(value)
			return true
		}
	}

	return false
}

func (s *detaCloud) GetUserWithCredentials(email, pwd string) (models.User, error) {
	var result *[]models.User

	db, err := base.New(s.client, "usuario")
	s.logError("failed init base instance", err)

	query := base.Query{
		{
			"email":    email,
			"password": pwd,
		},
	}

	_, err = db.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})
	s.logError("failed to fetch items", err)

	if len(*result) > 0 {
		for _, value := range *result {
			return value, nil
		}
	}

	return models.User{}, nil
}
func (s *detaCloud) InsertUser(u models.User) error {
	db, err := base.New(s.client, "usuario")
	s.logError("failed init base instance", err)

	log.Println(u)

	if ok := s.existsEmail(u.Email); ok {
		return errors.New("Email ya registrado")
	}

	key, err := db.Insert(u)
	if err != nil {
		s.logError("failed to insert item", err)
		return err
	}
	log.Println(key)

	return nil
}

// implement UserService interface
func (s *detaCloud) GetUsers() ([]models.User, error) {
	var result []models.User

	db, err := base.New(s.client, "usuario")
	s.logError("failed init base instance", err)

	_, err = db.Fetch(&base.FetchInput{
		Dest: &result,
	})
	s.logError("failed to fetch items", err)

	return result, nil
}
