package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageUserService interface {
	GetUserWithCredentials(email, pwd string) (models.User, error)
	//GetUsers() []*models.User
}
type StorageComputerService interface {
	GetComputers() []*models.CPU
}

type AuthService interface {
	AuthUser(email, pwd string) (models.User, error)
	ValidateTokenCookie(token string) bool
}

type ComputerService interface {
	ListComputers() []*models.CPU
}
