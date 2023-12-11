package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/database"
)

type DeviceRepository interface {
	Save(device *domain.Device) error
	Delete(name string) error
	Find(name string) (domain.Device, error)
	FindAll() ([]domain.Device, error)

	CollectorIsUsed(collector *domain.Collector) bool    // 是否引用了采集接口
	DeviceTypeIsUsed(deviceType *domain.DeviceType) bool // 是否引用了设备类型
	CollectorChanged(collector domain.Collector)         // 采集接口改变了，同步更新设备列表
	DeviceTypeChanged(deviceType domain.DeviceType)      // 设备类型改变了，同步更新设备列表
}

var _ DeviceRepository = (*deviceRepository)(nil)

type deviceRepository struct {
	db *bolt.DB
}

func NewDeviceRepository(db *bolt.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (u *deviceRepository) Save(device *domain.Device) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Device))

		if data, err := json.Marshal(device); err != nil {
			return err
		} else {
			return b.Put([]byte(device.Name), data)
		}
	})
}

func (u *deviceRepository) Delete(name string) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Device))

		return b.Delete([]byte(name))
	})
}

func (u *deviceRepository) Find(name string) (domain.Device, error) {
	device := domain.Device{}

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Device))
		data := b.Get([]byte(name))

		if data != nil {
			return json.Unmarshal(data, &device)
		} else {
			return errors.New("没有找到相关设备")
		}

	})

	return device, err
}

func (u *deviceRepository) FindAll() ([]domain.Device, error) {
	devices := make([]domain.Device, 0)

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Device))

		return b.ForEach(func(k, v []byte) error {
			device := domain.Device{}

			err := json.Unmarshal(v, &device)
			if err == nil {
				devices = append(devices, device)
			}
			return err
		})
	})

	return devices, err
}

func (u *deviceRepository) CollectorIsUsed(collector *domain.Collector) bool {
	if devices, err := u.FindAll(); err == nil {
		for _, device := range devices {
			if device.Collector.Name == collector.Name {
				return true
			}
		}
	}

	return false
}

func (u *deviceRepository) DeviceTypeIsUsed(deviceType *domain.DeviceType) bool {
	if devices, err := u.FindAll(); err == nil {
		for _, device := range devices {
			if device.Type.Name == deviceType.Name {
				return true
			}
		}
	}

	return false
}

func (u *deviceRepository) CollectorChanged(collector domain.Collector) {
	if devices, err := u.FindAll(); err == nil {
		for _, device := range devices {
			if device.Collector.Name == collector.Name {
				device.Collector = collector
				u.Save(&device)
			}
		}
	}
}

func (u *deviceRepository) DeviceTypeChanged(deviceType domain.DeviceType) {
	if devices, err := u.FindAll(); err == nil {
		for _, device := range devices {
			if device.Type.Name == deviceType.Name {
				device.Type = deviceType
				u.Save(&device)
			}
		}
	}
}
