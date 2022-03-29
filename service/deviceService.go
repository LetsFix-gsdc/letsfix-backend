package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"
)

type DeviceService interface {
	SaveDevice(models.Device) models.Device
	UpdateDevice(models.Device)
	DeleteDevice(models.Device)
	FindAllDevices() []models.Device
	FindDevice(uint) models.Device
}

type deviceService struct {
	deviceRepository repository.DeviceRepository
}

func NewDeviceService(repo repository.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepository: repo,
	} 
}

func (service *deviceService) SaveDevice(device models.Device) models.Device {
	service.deviceRepository.SaveDevice(device)
	return device
}

func (service *deviceService) FindAllDevices() []models.Device {
	return service.deviceRepository.FindAllDevices()
}

func (service *deviceService) UpdateDevice(device models.Device) {
	service.deviceRepository.UpdateDevice(device)
}

func (service *deviceService) DeleteDevice(device models.Device) {
	service.deviceRepository.DeleteDevice(device)
}

func (service *deviceService) FindDevice(device_id uint) models.Device {
	return service.deviceRepository.FindDeviceById(device_id)
}
