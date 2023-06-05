package models

type User struct {
	ID          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type UserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID          string `json:"id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
