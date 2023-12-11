package service

import (
	"errors"
	"smartgw/api/domain"
	"smartgw/api/repository"
)

type (
	DeviceTypeService interface {
		Add(deviceType *domain.DeviceType) error
		Update(deviceType *domain.DeviceType) error
		Delete(name string) error
		Find(name string) (domain.DeviceType, error)
		FindAll() ([]domain.DeviceType, error)

		AddProperties(deviceTypeName string, deviceProperty *domain.DeviceProperty) error
		UpdateProperties(deviceTypeName string, devicePropertyId int, deviceProperty *domain.DeviceProperty) error
		DeleteProperties(deviceTypeName string, devicePropertyId int) error
		FindProperty(deviceTypeName string, devicePropertyId int) (domain.DeviceProperty, error)
		FindAllProperties(deviceTypeName string) ([]domain.DeviceProperty, error)
	}

	deviceTypeService struct {
		deviceTypeRepository repository.DeviceTypeRepository
		deviceRepository     repository.DeviceRepository
	}
)

var _ DeviceTypeService = (*deviceTypeService)(nil)

func NewDeviceTypeService(deviceTypeRepository repository.DeviceTypeRepository, deviceRepository repository.DeviceRepository) DeviceTypeService {
	return &deviceTypeService{
		deviceTypeRepository: deviceTypeRepository,
		deviceRepository:     deviceRepository,
	}
}

func (dts *deviceTypeService) Add(deviceType *domain.DeviceType) error {
	if _, err := dts.deviceTypeRepository.Find(deviceType.Name); err == nil {
		return errors.New("设备类型已经存在！")
	}
	err := dts.deviceTypeRepository.Save(deviceType)
	if err != nil {
		return err
	}
	return err
}

func (dts *deviceTypeService) Update(deviceType *domain.DeviceType) error {
	err := dts.deviceTypeRepository.Save(deviceType)
	if err != nil {
		return err
	}
	dts.deviceRepository.DeviceTypeChanged(*deviceType)
	return err
}

func (dts *deviceTypeService) Delete(name string) error {
	// 检测设备是否引用了采集接口
	deviceType, err := dts.Find(name)
	if err == nil {
		if dts.deviceRepository.DeviceTypeIsUsed(&deviceType) {
			return errors.New("已经有设备引用了设备类型，不能删除！")
		}
	}

	err = dts.deviceTypeRepository.Delete(name)
	if err != nil {
		return err
	}
	return err
}

func (dts *deviceTypeService) Find(name string) (domain.DeviceType, error) {
	return dts.deviceTypeRepository.Find(name)
}

func (dts *deviceTypeService) FindAll() ([]domain.DeviceType, error) {
	return dts.deviceTypeRepository.FindAll()
}

func (dts *deviceTypeService) AddProperties(deviceTypeName string, deviceProperty *domain.DeviceProperty) error {
	err := dts.deviceTypeRepository.AddProperties(deviceTypeName, deviceProperty)
	if err != nil {
		return err
	}

	deviceType, _ := dts.Find(deviceTypeName)
	dts.deviceRepository.DeviceTypeChanged(deviceType)
	return err
}

func (dts *deviceTypeService) UpdateProperties(deviceTypeName string, devicePropertyId int, deviceProperty *domain.DeviceProperty) error {
	err := dts.deviceTypeRepository.UpdateProperties(deviceTypeName, devicePropertyId, deviceProperty)
	if err != nil {
		return err
	}

	deviceType, _ := dts.Find(deviceTypeName)
	dts.deviceRepository.DeviceTypeChanged(deviceType)
	return err
}

func (dts *deviceTypeService) DeleteProperties(deviceTypeName string, devicePropertyId int) error {
	err := dts.deviceTypeRepository.DeleteProperties(deviceTypeName, devicePropertyId)
	if err != nil {
		return err
	}

	deviceType, _ := dts.Find(deviceTypeName)
	dts.deviceRepository.DeviceTypeChanged(deviceType)
	return err
}

func (dts *deviceTypeService) FindProperty(deviceTypeName string, devicePropertyId int) (domain.DeviceProperty, error) {
	return dts.deviceTypeRepository.FindProperty(deviceTypeName, devicePropertyId)
}

func (dts *deviceTypeService) FindAllProperties(deviceTypeName string) ([]domain.DeviceProperty, error) {
	return dts.deviceTypeRepository.FindAllProperties(deviceTypeName)
}
