package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageNetworkService interface {
	SelectNetworks() []models.Network
	SelectResumeNetworks() []models.ResumeNetworks
}

type NetworkService interface {
	ListNetworks() []models.Network
	ListResumeNetworks() []models.ResumeNetworks
}
