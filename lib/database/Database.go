package database

import (
	"github.com/boltdb/bolt"
	"go.uber.org/zap"
)

const (
	Collector   = "Collector"   // 采集器
	DeviceType  = "DeviceType"  // 设备类型
	Device      = "Device"      // 设备
	CollectTask = "CollectTask" // 采集任务
	ReportTask  = "ReportTask"  // 上报任务
	User        = "User"        // 系统用户
)

func NewDatabase() (*bolt.DB, error) {
	db, err := bolt.Open("smart.db", 0600, nil)

	if err != nil {
		zap.L().Error("打开数据库失败！")
	} else {
		zap.L().Info("打开数据库成功！")
	}
	initDatabase(db)
	return db, err
}

func initDatabase(db *bolt.DB) {
	zap.L().Info("初始化数据库...")
	bucketNames := []string{Collector, DeviceType, Device, CollectTask, ReportTask, User}

	db.Update(func(tx *bolt.Tx) error {
		for _, name := range bucketNames {
			if _, err := tx.CreateBucketIfNotExists([]byte(name)); err != nil {
				return err
			}
		}
		return nil
	})
}
