package ports

import "github.com/luispfcanales/inventory-oti/models"

// services user
type StorageUserService interface {
	SelectUserWithCredentials(email, pwd string) (models.User, error)
	SelectUsers() ([]models.User, error)
	SelectStaff() ([]models.Staff, error)
	SelectRole() ([]models.Role, error)
	InsertUser(*models.User) error
}

type UserService interface {
	//GetByID(id string) (models.User, error)
	ListAllUsers() ([]models.User, error)
	SaveUser(*models.User) error
	ListStaff() ([]models.Staff, error)
	ListRole() ([]models.Role, error)
}
