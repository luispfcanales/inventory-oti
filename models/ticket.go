package models

import "encoding/json"

type Ticket struct {
	ID               string          `json:"id,omitempty"`
	IDServiceSupport string          `json:"id_service_support,omitempty"`
	IDStateTicket    string          `json:"id_state_ticket,omitempty"`
	CreatedByIDUser  string          `json:"created_by_id_user,omitempty"`
	CreatedAt        string          `json:"created_at,omitempty"`
	EndedAt          string          `json:"ended_at,omitempty"`
	MoreInfo         json.RawMessage `json:"more_info,omitempty"`
}
