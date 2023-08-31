package ports

import (
	"github.com/gofiber/websocket/v2"
	"github.com/luispfcanales/inventory-oti/models"
)

// services user
type StorageUserService interface {
	GetUserWithCredentials(email, pwd string) (models.User, error)
	GetUsers() ([]models.User, error)
	InsertUser(models.User) error
}
type AuthService interface {
	AuthUser(email, pwd string) (models.User, error)
	ValidateTokenCookie(token string) bool
}
type UserService interface {
	//GetByID(id string) (models.User, error)
	ListAll() ([]models.User, error)
	Save(models.User) error
}

// services computers
type StorageComputerService interface {
	GetComputers() []*models.CPU
}
type ComputerService interface {
	ListComputers() []*models.CPU
}

type StramingComputerService interface {
	AddConnection(id, role string, c *websocket.Conn)
	RemoveConnection(id, role string)
	Broadcast(*models.StreamEvent)
	ListAllConnections(chan<- models.StreamEvent)
	Receiver(func())
}

// api service consumer
type ApiService interface {
	GetDataByDni(dni string) (models.PersonServiceDNI, error)
}
