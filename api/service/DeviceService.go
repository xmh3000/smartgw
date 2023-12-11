package service

import (
	"errors"
	"smartgw/api/domain"
	"smartgw/api/repository"
)

type (
	DeviceService interface {
		Add(device *domain.Device) error
		Update(device *domain.Device) error
		Delete(name string) error
		Find(name string) (domain.Device, error)
		FindAll() ([]domain.Device, error)
	}

	deviceService struct {
		deviceRepository repository.DeviceRepository
	}
)

var _ DeviceService = (*deviceService)(nil)

func NewDeviceService(deviceRepository repository.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepository: deviceRepository,
	}
}

func (s *deviceService) Add(device *domain.Device) error {
	if _, err := s.deviceRepository.Find(device.Name); err == nil {
		return errors.New("设备已经存在！")
	}
	return s.deviceRepository.Save(device)
}

func (s *deviceService) Update(device *domain.Device) error {
	return s.deviceRepository.Save(device)
}

func (s *deviceService) Delete(name string) error {
	return s.deviceRepository.Delete(name)
}

func (s *deviceService) Find(name string) (domain.Device, error) {
	return s.deviceRepository.Find(name)
}

func (s *deviceService) FindAll() ([]domain.Device, error) {
	return s.deviceRepository.FindAll()
}
