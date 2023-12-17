package models

type Zone struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	FloorNumbers int    `json:"floor_numbers,omitempty"`
	Color        string `json:"color,omitempty"`
	IDCampus     string `json:"id_campus,omitempty"`
	NameCampus   string `json:"name_campus,omitempty"`
}

type Campus struct {
	ID           string `json:"id,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Name         string `json:"name,omitempty"`
	Address      string `json:"address,omitempty"`
	State        bool   `json:"state"`
}
