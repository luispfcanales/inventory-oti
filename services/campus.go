package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type CampusSrv struct {
	repo ports.StorageCampusService
}

func (campussrv *CampusSrv) PutCampus(c models.Campus) (models.Campus, error) {
	data, err := campussrv.repo.UpdateCampus(c)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (campussrv *CampusSrv) ListCampus(id string) (models.Campus, error) {
	data, err := campussrv.repo.SelectCampus(id)
	if err != nil {
		return data, errors.New("[ No se encontraron datos ]")
	}
	return data, nil
}

func (campussrv *CampusSrv) ListAllCampus() []models.Campus {
	return campussrv.repo.SelectAllCampus()
}

func (campussrv *CampusSrv) SaveCampus(c *models.Campus) error {
	c.ID = uuid.New().String()

	err := campussrv.repo.InsertCampus(*c)
	if err != nil {
		return err
	}

	return nil
}

func (campussrv *CampusSrv) RemoveCampus(key string) error {
	return campussrv.repo.DeleteCampus(key)
}

func NewCampus(r ports.StorageCampusService) ports.CampusService {
	return &CampusSrv{
		repo: r,
	}
}
