package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/config"
	"smartgw/lib/database"
	"strconv"
)

type ReportTaskRepository interface {
	Save(reportTask *domain.ReportTask) error
	Delete(name string) error
	Find(name string) (domain.ReportTask, error)
	FindAll() ([]domain.ReportTask, error)
	Migrate() error
}

var _ ReportTaskRepository = (*reportTaskRepository)(nil)

type reportTaskRepository struct {
	db     *bolt.DB
	config *config.Config
}

func NewReportTaskRepository(db *bolt.DB, config *config.Config) ReportTaskRepository {
	return &reportTaskRepository{
		db,
		config,
	}
}

func (rt *reportTaskRepository) Migrate() error {
	port, _ := strconv.Atoi(rt.config.Gateway.Port)
	task, err := rt.Find("数据上报")
	// 新增
	if err != nil {
		task.Name = "数据上报"
		task.ReportName = "物联平台"
		task.Ip = rt.config.Ip
		task.Port = port
		task.ClientID = rt.config.ClientID
		task.Username = rt.config.Username
		task.Password = rt.config.Password
		task.Cron = "@every 1m30s"
		task.Status = 0
	} else {
		task.Ip = rt.config.Ip
		task.Port = port
		task.ClientID = rt.config.ClientID
		task.Username = rt.config.Username
		task.Password = rt.config.Password
	}
	//return rt.Save(&task)

	return rt.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.ReportTask))

		if data, err := json.Marshal(&task); err != nil {
			return err
		} else {
			return b.Put([]byte(task.Name), data)
		}
	})
}

func (rt *reportTaskRepository) Save(reportTask *domain.ReportTask) error {

	rt.config.Gateway.Ip = reportTask.Ip
	rt.config.Gateway.Port = strconv.Itoa(reportTask.Port)
	rt.config.Gateway.ClientID = reportTask.ClientID

	config.SaveConfig(rt.config)

	return rt.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.ReportTask))

		if data, err := json.Marshal(reportTask); err != nil {
			return err
		} else {
			return b.Put([]byte(reportTask.Name), data)
		}
	})
}

func (rt *reportTaskRepository) Delete(name string) error {
	return rt.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.ReportTask))

		return b.Delete([]byte(name))
	})
}

func (rt *reportTaskRepository) Find(name string) (domain.ReportTask, error) {
	reportTask := domain.ReportTask{}

	err := rt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.ReportTask))
		data := b.Get([]byte(name))

		if data != nil {
			return json.Unmarshal(data, &reportTask)
		} else {
			return errors.New("没有找到相关数据")
		}

	})
	return reportTask, err
}

func (rt *reportTaskRepository) FindAll() ([]domain.ReportTask, error) {
	reportTasks := make([]domain.ReportTask, 0)

	err := rt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.ReportTask))

		return b.ForEach(func(k, v []byte) error {
			reportTask := domain.ReportTask{}
			err := json.Unmarshal(v, &reportTask)
			if err == nil {
				reportTasks = append(reportTasks, reportTask)
			}
			return err
		})
	})

	return reportTasks, err
}
