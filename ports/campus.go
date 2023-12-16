package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageCampusService interface {
	InsertCampus(models.Campus) error
	DeleteCampus(string) error
	UpdateCampus(models.Campus) (models.Campus, error)
	SelectCampus(string) (models.Campus, error)
	SelectAllCampus() []models.Campus
}

type CampusService interface {
	SaveCampus(*models.Campus) error
	RemoveCampus(string) error
	PutCampus(models.Campus) (models.Campus, error)
	ListCampus(string) (models.Campus, error)
	ListAllCampus() []models.Campus
}
