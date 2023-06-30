package clouddeta

import (
	"log"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/luispfcanales/inventory-oti/models"
)

type detaCloud struct {
	client *deta.Deta
}

func NewCloudDetaStorage(key string) *detaCloud {
	d, err := deta.New(deta.WithProjectKey(key))
	if err != nil {
		log.Fatal("aqui", err)
	}

	cloud := &detaCloud{
		client: d,
	}

	return cloud
}

func (s *detaCloud) GetComputers() []*models.CPU {
	var result []*models.CPU

	db, err := base.New(s.client, "equipos")
	s.logError("failed init base instance", err)

	_, err = db.Fetch(&base.FetchInput{
		Dest: &result,
	})
	s.logError("failed to fetch items", err)

	return result
}

func (s *detaCloud) GetUserWithCredentials(email, pwd string) (models.User, error) {
	var result *[]models.User

	db, err := base.New(s.client, "usuario")
	s.logError("failed init base instance", err)

	query := base.Query{
		{
			"Email":    email,
			"Password": pwd,
		},
	}

	_, err = db.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &result,
	})
	s.logError("failed to fetch items", err)

	if len(*result) > 0 {
		for _, value := range *result {
			return value, nil
		}
	}

	return models.User{}, nil
}

func (s *detaCloud) logError(text string, err error) {
	if err != nil {
		log.Fatalf("STORAGE: %s -> %v :", text, err)
	}
}
