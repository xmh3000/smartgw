package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/database"
)

type CollectorRepository interface {
	Save(collector *domain.Collector) error
	Delete(name string) error
	Find(name string) (domain.Collector, error)
	FindAll() ([]domain.Collector, error)
}

var _ CollectorRepository = (*collectorRepository)(nil)

type collectorRepository struct {
	db *bolt.DB
}

func NewCollectorRepository(db *bolt.DB) CollectorRepository {
	return &collectorRepository{
		db: db,
	}
}

func (c *collectorRepository) Save(collector *domain.Collector) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Collector))

		if data, err := json.Marshal(collector); err != nil {
			return err
		} else {
			return b.Put([]byte(collector.Name), data)
		}
	})
}

func (c *collectorRepository) Delete(name string) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Collector))

		return b.Delete([]byte(name))
	})
}

func (c *collectorRepository) Find(name string) (domain.Collector, error) {
	collector := domain.Collector{}

	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Collector))
		data := b.Get([]byte(name))
		if data != nil {
			return json.Unmarshal(data, &collector)
		} else {
			return errors.New("没有找到相关采集设备")
		}
	})

	return collector, err
}

func (c *collectorRepository) FindAll() ([]domain.Collector, error) {
	collectors := make([]domain.Collector, 0)

	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.Collector))

		return b.ForEach(func(k, v []byte) error {
			collector := domain.Collector{}
			err := json.Unmarshal(v, &collector)
			if err == nil {
				collectors = append(collectors, collector)
			}
			return err
		})
	})

	return collectors, err
}
