package services

import (
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type DeviceSrv struct {
	repo ports.StorageDeviceService
}

func (devicesrv *DeviceSrv) SaveDevice(_ models.Device) (models.Device, error) {
	panic("not implemented") // TODO: Implement
}

func (devicesrv *DeviceSrv) ListAllDevice() []models.Device {
	return devicesrv.repo.SelectAllDevice()
}

func NewDevice(r ports.StorageDeviceService) ports.DeviceService {
	return &DeviceSrv{
		repo: r,
	}
}
