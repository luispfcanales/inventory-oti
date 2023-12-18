package models

import (
	"encoding/json"
)

type Device struct {
	ID               int             `json:"id,omitempty"`
	PatrimonialCode  string          `json:"patrimonial_code,omitempty"`
	SerialCode       string          `json:"serial_code,omitempty"`
	Brand            string          `json:"brand,omitempty"`
	ModelName        string          `json:"model_name,omitempty"`
	StateDevice      string          `json:"state_device,omitempty"`
	TypeDevice       string          `json:"type_device,omitempty"`
	DependencyDevice string          `json:"dependency_device,omitempty"`
	AdquisitonDate   string          `json:"adquisiton_date,omitempty"`
	MoreInfo         json.RawMessage `json:"more_info,omitempty"`
}
