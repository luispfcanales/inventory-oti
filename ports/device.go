package ports

import "github.com/luispfcanales/inventory-oti/models"

type StorageDeviceService interface {
	InsertDevice(models.Device) (models.Device, error)
	SelectAllDevice() []models.Device
}

type DeviceService interface {
	SaveDevice(models.Device) (models.Device, error)
	ListAllDevice() []models.Device
}
