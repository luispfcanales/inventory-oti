package services

import (
	"github.com/luispfcanales/inventory-oti/models"
	"github.com/luispfcanales/inventory-oti/ports"
)

type DeviceSrv struct {
	repo ports.StorageDeviceService
}

func (devicesrv *DeviceSrv) ListAllDevice() []models.Device {
	return devicesrv.repo.SelectAllDevice()
}

func NewDevice(r ports.StorageDeviceService) ports.DeviceService {
	return &DeviceSrv{
		repo: r,
	}
}
