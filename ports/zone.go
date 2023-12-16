package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageZoneService interface {
	SelectAllZone() []models.Zone
	SelectZone(string) (models.Zone, error)
}
type ZoneService interface {
	ListAllZone() []models.Zone
	ListZone(string) (models.Zone, error)
}
