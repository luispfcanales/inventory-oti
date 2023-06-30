package models

type CPU struct {
	Key          string `json:"key,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Disk         string `json:"disk,omitempty"`
	Maker        string `json:"maker,omitempty"`
	Model        string `json:"model,omitempty"`
	Name         string `json:"name,omitempty"`
	Processor    string `json:"processor,omitempty"`
	Ram          string `json:"ram,omitempty"`
	Serial       string `json:"serial,omitempty"`
	SizeDisk     string `json:"size_disk,omitempty"`
}
