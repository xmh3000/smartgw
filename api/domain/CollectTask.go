package domain

// CollectTask 采集任务，目前有且只有一个采集任务，采集周期设置 2.5 * 最大接入设备数（多个通道时，只计算最大的设备数）
type CollectTask struct {
	Name   string `json:"name"`   //名称
	Cron   string `json:"cron"`   //定时策略
	Status int    `json:"status"` //状态
}
