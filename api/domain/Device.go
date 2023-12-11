package domain

import "time"

// Device 设备列表
type Device struct {
	Name      string     `json:"name"`      //设备名称
	Type      DeviceType `json:"type"`      //设备类型
	Address   string     `json:"address"`   //通讯地址
	Collector Collector  `json:"collector"` //采集接口
	// Alone、Serial 可根据实际情况去掉，即一个采集接口有【固定】的波特率和驱动程序-----------
	Alone  bool   `json:"alone"`  //独立开关
	Serial Serial `json:"serial"` //通讯参数
	//----------------------------------------------------------------------------
	Online         bool      `json:"online"`         //设备在线
	CollectTime    time.Time `json:"collectTime"`    //最后采集时间
	CollectTotal   int       `json:"collectTotal"`   //总采集次数
	CollectSuccess int       `json:"collectSuccess"` //采集成功次数
	//----------------------------------------------------------------------------
	ReportTime    time.Time `json:"reportTime"`    //最后上报时间
	ReportTotal   int       `json:"reportTotal"`   //总上报次数
	ReportSuccess int       `json:"reportSuccess"` //上报成功次数
}
