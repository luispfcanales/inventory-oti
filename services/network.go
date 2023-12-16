package services

import (
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type NetworkSrv struct {
	repo ports.StorageNetworkService
}

func (networksrv *NetworkSrv) ListResumeNetworks() []models.ResumeNetworks {
	return networksrv.repo.SelectResumeNetworks()
}

func (networksrv *NetworkSrv) ListNetworks() []models.Network {
	return networksrv.repo.SelectNetworks()
}

func NewNetwork(r ports.StorageNetworkService) ports.NetworkService {
	return &NetworkSrv{
		repo: r,
	}
}
