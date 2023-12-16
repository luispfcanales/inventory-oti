package services

import (
	"errors"

	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type UserSrv struct {
	repo ports.StorageUserService
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

func (us *UserSrv) SaveUser(u models.User) error {
	err := us.repo.InsertUser(u)
	if err != nil {
		return err
	}
	return nil
}
