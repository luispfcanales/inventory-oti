package services

import (
	"errors"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type PersonSrv struct {
	repo ports.StoragePersonService
}

func (personsrv *PersonSrv) RemovePerson(dni int) bool {
	if !personsrv.repo.DeletePerson(dni) {
		return false
	}
	return true
}

func (personsrv *PersonSrv) SavePerson(p models.Person) (models.Person, error) {
	data, err := personsrv.repo.InsertPerson(p)
	if err != nil {
		log.Println(err)
		return p, errors.New("Person not insert in storage")
	}
	return data, nil
}

func (personsrv *PersonSrv) ListOne(dni int) (models.Person, error) {
	p, err := personsrv.repo.SelectPerson(dni)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (personsrv *PersonSrv) ListAll() []models.Person {
	return personsrv.repo.SelectPersons()
}

func (personsrv *PersonSrv) Put(p models.Person) (models.Person, error) {
	p, err := personsrv.repo.UpdatePerson(p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func NewPerson(r ports.StoragePersonService) ports.PersonService {
	return &PersonSrv{
		repo: r,
	}
}
