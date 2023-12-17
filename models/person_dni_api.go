package models

type PersonServiceDNI struct {
	Success        bool   `json:"success,omitempty"`
	Status         string `json:"status,omitempty"`
	Message        string `json:"message,omitempty"`
	ErrorCode      int    `json:"error_code,omitempty"`
	Name           string `json:"nombres,omitempty"`
	FatherLastName string `json:"apellidoPaterno,omitempty"`
	MotherLastName string `json:"apellidoMaterno,omitempty"`
}
