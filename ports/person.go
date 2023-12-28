package ports

import "github.com/luispfcanales/inventory-oti/models"

// Person -select-update-delete-insert
type StoragePersonService interface {
	DeletePerson(dni int) bool
	InsertPerson(models.Person) (models.Person, error)
	SelectPerson(dni string) (models.Person, error)
	SelectPersons() []models.Person
	UpdatePerson(models.Person) (models.Person, error)
}

// list-put-remove-save
type PersonService interface {
	RemovePerson(int) bool
	SavePerson(models.Person) (models.Person, error)
	ListOne(string) (models.Person, error)
	ListAll() []models.Person
	Put(models.Person) (models.Person, error)
}
