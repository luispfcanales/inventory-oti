package models

type OptionsSignature struct {
	StampAppearanceID string `json:"stamp_appearance_id,omitempty"`
	FileID            string `json:"file_id,omitempty"`
	PageNumber        string `json:"page_number,omitempty"`
	Pox               string `json:"pox,omitempty"`
	Poy               string `json:"poy,omitempty"`
	Reason            string `json:"reason,omitempty"`
}
