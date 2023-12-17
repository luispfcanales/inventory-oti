package models

import "time"

type Person struct {
	IDPerson  int       `json:"dni,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Address   string    `json:"address,omitempty"`
}

type User struct {
	Key      string `json:"key,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	Active   bool   `json:"active,omitempty"`
	IDRole   string `json:"id_role,omitempty"`
	Staff    string `json:"staff,omitempty"`
	Person

	AccessToken string `json:"access_token,omitempty"`
}

type UserRequest struct {
	Username string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID          int    `json:"id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	Username    string `json:"username,omitempty"`
	Role        string `json:"role,omitempty"`
	Staff       string `json:"staff,omitempty"`
	Active      bool   `json:"active"`
}
