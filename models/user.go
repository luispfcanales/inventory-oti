package models

type User struct {
	Key       string `json:"key,omitempty"`
	Dni       string `json:"dni,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	State     int    `json:"state"`
	IDRole    string `json:"id_role,omitempty"`
	IDStaff   string `json:"id_staff,omitempty"`
	Email     string `json:"email,omitempty"`

	AccessToken string `json:"access_token,omitempty"`
}

type UserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID          string `json:"id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	Username    string `json:"username,omitempty"`
}
