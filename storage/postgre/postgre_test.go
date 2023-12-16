package postgre

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luispfcanales/inventory-oti/models"
)

func TestConnectionDB(t *testing.T) {
	stg := NewPostgreStorage()
	err := stg.getConnection().Ping()
	if err != nil {
		t.Fail()
	}
}

func TestGetUserWithCredentials(t *testing.T) {
	stg := NewPostgreStorage()
	dni_want := 72453560
	got, err := stg.SelectUserWithCredentials(
		"luispfcanales@gmail.com",
		"1234",
	)

	if err != nil {
		t.Fail()
	}

	if got.IDPerson != dni_want {
		t.Fail()
	}
}

func TestGetUsers(t *testing.T) {
	stg := NewPostgreStorage()
	got, err := stg.SelectUsers()
	if err != nil {
		t.Fail()
	}

	if len(got) == 0 {
		t.Fail()
	}
}

// tests port StoragePersonService
func TestUpdatePerson(t *testing.T) {
	stg := NewPostgreStorage()

	want := models.Person{
		IDPerson:  72453560,
		FirstName: "LUIS ANGEL",
		LastName:  "PFUÑO CANALES",
		Birthdate: time.Now(),
		Address:   "JR ANCASH 865",
	}

	got, err := stg.UpdatePerson(want)
	if err != nil {
		t.Fail()
	}

	if got.IDPerson != want.IDPerson {
		t.Fail()
	}
}

func TestSelectPersons(t *testing.T) {
	stg := NewPostgreStorage()

	got := stg.SelectPersons()
	if len(got) != 2 {
		t.Fail()
	}
}

func TestSelectPerson(t *testing.T) {
	stg := NewPostgreStorage()

	want := 72453560

	got, err := stg.SelectPerson(want)
	if err != nil {
		t.Fail()
	}

	if got.IDPerson != want {
		t.Fail()
	}
}

func TestInsertPerson(t *testing.T) {
	stg := NewPostgreStorage()

	want := models.Person{
		IDPerson:  72453559,
		FirstName: "prueba",
		LastName:  "prueba",
		Birthdate: time.Now(),
		Address:   "jr test",
	}

	got, err := stg.InsertPerson(want)
	if err != nil {
		t.Fail()
	}

	if got.IDPerson != want.IDPerson {
		t.Fail()
	}
}

func TestDeletePerson(t *testing.T) {
	stg := NewPostgreStorage()

	ok := stg.DeletePerson(72453559)
	if !ok {
		t.Fail()
	}
}

// test networks
func TestSelectNetworks(t *testing.T) {
	stg := NewPostgreStorage()

	want := 49 //registered ip address
	got := stg.SelectNetworks()
	if len(got) != want {
		t.Fail()
	}

}

// test campus
func TestInsertAndDeleteCampus(t *testing.T) {
	stg := NewPostgreStorage()

	key := uuid.New().String()
	want := models.Campus{
		ID:           key,
		Abbreviation: "TEST",
		Name:         "Testing golang",
		Address:      "golang.com",
		State:        true,
	}

	err := stg.InsertCampus(want)
	if err != nil {
		t.Fail()
	}

	err = stg.DeleteCampus(key)
	if err != nil {
		t.Fail()
	}
}
