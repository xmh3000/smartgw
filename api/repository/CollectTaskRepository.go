package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/database"
)

type CollectTaskRepository interface {
	Save(collectTask *domain.CollectTask) error
	Delete(name string) error
	Find(name string) (domain.CollectTask, error)
	FindAll() ([]domain.CollectTask, error)
	Migrate() error
}

var _ CollectTaskRepository = (*collectTaskRepository)(nil)

type collectTaskRepository struct {
	db *bolt.DB
}

var _ CollectTaskRepository = (*collectTaskRepository)(nil)

func NewCollectTaskRepository(db *bolt.DB) CollectTaskRepository {
	return &collectTaskRepository{
		db: db,
	}
}

func (rt *collectTaskRepository) Migrate() error {
	task, err := rt.Find("数据采集")
	// 新增
	if err != nil {
		task.Name = "数据采集"
		task.Cron = "@every 1m30s"
		task.Status = 0
	}
	return rt.Save(&task)
}

func (u *collectTaskRepository) Save(collectTask *domain.CollectTask) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.CollectTask))

		if data, err := json.Marshal(collectTask); err != nil {
			return err
		} else {
			return b.Put([]byte(collectTask.Name), data)
		}
	})
}

func (u *collectTaskRepository) Delete(name string) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.CollectTask))

		return b.Delete([]byte(name))
	})
}

func (u *collectTaskRepository) Find(name string) (domain.CollectTask, error) {
	collectTask := domain.CollectTask{}

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.CollectTask))
		data := b.Get([]byte(name))

		if data != nil {
			return json.Unmarshal(data, &collectTask)
		} else {
			return errors.New("没有找到相关采集任务")
		}
	})

	return collectTask, err
}

func (u *collectTaskRepository) FindAll() ([]domain.CollectTask, error) {
	collectTasks := make([]domain.CollectTask, 0)

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.CollectTask))

		return b.ForEach(func(k, v []byte) error {
			collectTask := domain.CollectTask{}
			err := json.Unmarshal(v, &collectTask)
			if err == nil {
				collectTasks = append(collectTasks, collectTask)
			}
			return err
		})
	})

	return collectTasks, err
}
