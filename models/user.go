package models

type User struct {
	Key      string `json:"key,omitempty"`
	Fullname string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	State    string `json:"state,omitempty"`
	IDRole   string `json:"id_role,omitempty"`
	IDStaff  string `json:"id_staff,omitempty"`
	Email    string `json:"email,omitempty"`

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
