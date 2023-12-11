package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/database"
)

type DeviceTypeRepository interface {
	// 对设备类型大类操作
	Save(deviceType *domain.DeviceType) error
	Delete(name string) error
	Find(name string) (domain.DeviceType, error)
	FindAll() ([]domain.DeviceType, error)
	// 对设备类型中的相关属性进行操作
	AddProperties(deviceTypeName string, deviceProperty *domain.DeviceProperty) error
	UpdateProperties(deviceTypeName string, devicePropertyId int, deviceProperty *domain.DeviceProperty) error
	DeleteProperties(deviceTypeName string, devicePropertyId int) error
	FindProperty(deviceTypeName string, devicePropertyId int) (domain.DeviceProperty, error)
	FindAllProperties(deviceTypeName string) ([]domain.DeviceProperty, error)
}

var _ DeviceTypeRepository = (*deviceTypeRepository)(nil)

type deviceTypeRepository struct {
	db *bolt.DB
}

func NewDeviceTypeRepository(db *bolt.DB) DeviceTypeRepository {
	return &deviceTypeRepository{
		db: db,
	}
}

func (u *deviceTypeRepository) Save(deviceType *domain.DeviceType) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.DeviceType))

		if data, err := json.Marshal(deviceType); err != nil {
			return err
		} else {
			return b.Put([]byte(deviceType.Name), data)
		}
	})
}

func (u *deviceTypeRepository) Delete(name string) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.DeviceType))

		return b.Delete([]byte(name))
	})
}

func (u *deviceTypeRepository) Find(name string) (domain.DeviceType, error) {
	deviceType := domain.DeviceType{}

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.DeviceType))
		data := b.Get([]byte(name))

		if data != nil {
			return json.Unmarshal(data, &deviceType)
		} else {
			return errors.New("没有找到相关设备类型")
		}

	})

	return deviceType, err
}

func (u *deviceTypeRepository) FindAll() ([]domain.DeviceType, error) {
	deviceTypes := make([]domain.DeviceType, 0)

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.DeviceType))

		return b.ForEach(func(k, v []byte) error {
			deviceType := domain.DeviceType{}
			err := json.Unmarshal(v, &deviceType)
			if err == nil {
				deviceTypes = append(deviceTypes, deviceType)
			}
			return err
		})
	})

	return deviceTypes, err
}

func (u *deviceTypeRepository) AddProperties(deviceTypeName string, deviceProperty *domain.DeviceProperty) error {
	deviceType, err := u.Find(deviceTypeName)
	if err != nil {
		return err
	}

	deviceType.Properties = append(deviceType.Properties, *deviceProperty)

	return u.Save(&deviceType)
}

func (u *deviceTypeRepository) UpdateProperties(deviceTypeName string, devicePropertyId int, deviceProperty *domain.DeviceProperty) error {
	deviceType, err := u.Find(deviceTypeName)
	if err != nil {
		return err
	}
	deviceType.Properties[devicePropertyId-1] = *deviceProperty

	return u.Save(&deviceType)
}

func (u *deviceTypeRepository) DeleteProperties(deviceTypeName string, devicePropertyId int) error {
	deviceType, err := u.Find(deviceTypeName)
	if err != nil {
		return err
	}

	deviceType.Properties = append(deviceType.Properties[:devicePropertyId-1], deviceType.Properties[devicePropertyId:]...)
	return u.Save(&deviceType)
}

func (u *deviceTypeRepository) FindProperty(deviceTypeName string, devicePropertyId int) (domain.DeviceProperty, error) {
	deviceType, err := u.Find(deviceTypeName)

	return deviceType.Properties[devicePropertyId-1], err
}

func (u *deviceTypeRepository) FindAllProperties(deviceTypeName string) ([]domain.DeviceProperty, error) {
	deviceType, err := u.Find(deviceTypeName)

	return deviceType.Properties, err
}
