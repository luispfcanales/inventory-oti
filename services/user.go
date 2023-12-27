package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type UserSrv struct {
	repo ports.StorageUserService
}

func (us *UserSrv) ListStaff() ([]models.Staff, error) {
	data, err := us.repo.SelectStaff()
	if err != nil {
		return nil, errors.New("[fn ListStaff] StorageService Error")
	}
	return data, nil
}
func (us *UserSrv) ListRole() ([]models.Role, error) {
	data, err := us.repo.SelectRole()
	if err != nil {
		return nil, errors.New("[fn ListRole] StorageService Error")
	}
	return data, nil
}

func NewUser(r ports.StorageUserService) ports.UserService {
	return &UserSrv{
		repo: r,
	}
}

func (us *UserSrv) ListAllUsers() ([]models.User, error) {
	listUsers, err := us.repo.SelectUsers()
	if err != nil {
		return nil, errors.New("[fn ListAll] StorageService Error")
	}
	return listUsers, nil
}

func (us *UserSrv) SaveUser(u *models.User) error {
	u.Key = uuid.New().String()
	err := us.repo.InsertUser(u)
	if err != nil {
		return err
	}
	return nil
}
