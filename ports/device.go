package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageDeviceService interface {
	SelectAllDevice() []models.Device
}

type DeviceService interface {
	ListAllDevice() []models.Device
}
