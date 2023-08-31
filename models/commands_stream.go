package models

type StreamEvent struct {
	ID            string      `json:"id,omitempty"`
	Status        string      `json:"status,omitempty"`
	Event         string      `json:"event,omitempty"`
	Role          string      `json:"role,omitempty"`
	EventEmisorID string      `json:"event_emisor_id,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
}
