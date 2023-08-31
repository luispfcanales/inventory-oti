package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type ApiSrv struct {
	token string
}

func loadToken() string {
	return os.Getenv("tokenApiDni")
}

func NewApiDni() ports.ApiService {
	return &ApiSrv{
		token: loadToken(),
	}
}

// GetDataByDni return info person
func (a *ApiSrv) GetDataByDni(dni string) (models.PersonServiceDNI, error) {
	//API_URL := "https://api-dni-ruc.vercel.app/jp_api/reniec?dni=%s"
	API_URL := "https://dniruc.apisperu.com/api/v1/dni/%s?token=%s"
	API_URL = fmt.Sprintf(API_URL, dni, a.token)

	data := models.PersonServiceDNI{}

	r, err := http.Get(API_URL)
	if err != nil {
		log.Println("[error making GET request]: ", err)
		return data, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[error reading request]: ", err)
		return data, err
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("[error parsing JSON]: ", err)
		return data, err
	}

	log.Println(data)
	return data, nil
}
