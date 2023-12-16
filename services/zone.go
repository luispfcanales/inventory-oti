package services

import (
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type ZoneSrv struct {
	repo ports.StorageZoneService
}

func (zonesrv *ZoneSrv) ListAllZone() []models.Zone {
	return zonesrv.repo.SelectAllZone()
}

func (zonesrv *ZoneSrv) ListZone(id string) (models.Zone, error) {
	data, err := zonesrv.repo.SelectZone(id)
	if err != nil {
		return models.Zone{}, err
	}
	return data, nil
}

func NewZone(r ports.StorageZoneService) ports.ZoneService {
	return &ZoneSrv{
		repo: r,
	}
}
