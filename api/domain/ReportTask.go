package domain

// ReportTask 上报任务，有且只有一条记录，对应的数据应该上报到物联平台（thingsped）
type ReportTask struct {
	Name       string `json:"name"`       // 任务名称
	ReportName string `json:"reportName"` // 平台名称
	Ip         string `json:"ip"`         // ip地址
	Port       int    `json:"port"`       // 端口
	ClientID   string `json:"clientID"`   // 网关编号
	Username   string `json:"username"`   // 用户名
	Password   string `json:"password"`   // 密码
	Cron       string `json:"cron"`       // cron表达式
	Status     int    `json:"status"`     // 运行状态 0：停止 1：运行
}
