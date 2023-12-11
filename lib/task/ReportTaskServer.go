package task

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"smartgw/api/domain"
	"smartgw/api/repository"
	"smartgw/lib/mqtt"
	"sync"
	"time"
)

type (
	// ReportTaskServer 定时上报任务
	ReportTaskServer struct {
		tasks            *sync.Map
		deviceRepository repository.DeviceRepository
	}
)

// NewReportTaskServer 实例化
func NewReportTaskServer(deviceRepository repository.DeviceRepository) *ReportTaskServer {
	return &ReportTaskServer{
		tasks:            &sync.Map{},
		deviceRepository: deviceRepository,
	}
}

// InitReportTaskServer 初始化
func InitReportTaskServer(reportTaskServer *ReportTaskServer, repository repository.ReportTaskRepository) {
	zap.S().Info("初始化上报定时任务")
	if tasks, err := repository.FindAll(); err == nil {
		for _, task := range tasks {
			reportTaskServer.Add(&task)
		}
	}
}

// Add 新增定时任务，如果定时任务是开启状态，则开启定时任务
func (rts *ReportTaskServer) Add(task *domain.ReportTask) {
	zap.S().Info("新增定时任务", task)
	c := cron.New()
	c.AddFunc(task.Cron, rts.Send)
	rts.tasks.Store(task.Name, c)

	// 启动定时任务
	if task.Status == 1 {
		rts.Start(task)
	}
}

// Delete 删除定时任务，如果定时任务是开启状态，则停止定时任务
func (rts *ReportTaskServer) Delete(name string) {
	zap.S().Info("删除定时任务", name)
	if value, loaded := rts.tasks.LoadAndDelete(name); loaded {
		c := value.(*cron.Cron)
		c.Stop()
	}
}

// Update 修改定时任务，删除存在的定时任务，同时增加一个新的定时任务
func (rts *ReportTaskServer) Update(task *domain.ReportTask) {
	rts.Delete(task.Name)
	rts.Add(task)
}

// Start 启动定时任务
func (rts *ReportTaskServer) Start(task *domain.ReportTask) {
	zap.S().Info("启动定时任务", task)
	if value, loaded := rts.tasks.Load(task.Name); loaded {
		c := value.(*cron.Cron)
		c.Start()
	}
}

// Stop 停止定时任务
func (rts *ReportTaskServer) Stop(task *domain.ReportTask) {
	zap.S().Info("停止定时任务", task)
	if value, loaded := rts.tasks.Load(task.Name); loaded {
		c := value.(*cron.Cron)
		c.Stop()
	}
}

// Send 数据上报
func (rts *ReportTaskServer) Send() {
	zap.S().Info("开始数据上报...")

	// 获取设备数据
	devices, _ := rts.deviceRepository.FindAll()

	if len(devices) < 1 {
		zap.S().Error("设备列表为空！")
		return
	}

	// 组报文，格式如下：
	// {
	//    "Device A" [{
	//       "ts": 2323423423324
	//       "values" [{},{}]
	nodes := make(map[string]interface{})
	for _, device := range devices {
		properties := make(map[string]interface{})

		// 如果设备在线，则上报采集数据
		for _, property := range device.Type.Properties {
			if property.Reported {
				properties[property.Name] = property.Value
			}
		}
		properties["ONLINE"] = device.Online
		telemetry := make(map[string]interface{})
		telemetry["ts"] = time.Now().Unix() * 1000
		telemetry["values"] = properties

		node := make([]map[string]interface{}, 1)
		node[0] = telemetry
		nodes[device.Name] = node
	}

	zap.S().Info("上报内容：", nodes)
	result, _ := json.Marshal(nodes)
	// 数据上报
	mqtt.MClient.Publish(result)
	zap.S().Info("上报成功!")
}
