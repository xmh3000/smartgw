package domain

// DeviceProperty 设备属性
type DeviceProperty struct {
	Name        string      `json:"name"`        //名称，英文
	Description string      `json:"description"` //描述，中文涵义
	Type        string      `json:"type"`        //类型:int,long,double,string
	Length      int         `json:"length"`      //长度
	Decimal     int         `json:"decimal"`     //小数位
	Unit        string      `json:"unit"`        //计量单位
	Value       interface{} `json:"value"`       //数值
	Reported    bool        `json:"reported"`    //是否上报
}

// DeviceType 设备类型
type DeviceType struct {
	Name       string           `json:"name"`
	Driver     string           `json:"driver"`     //驱动程序目录、文件名
	Properties []DeviceProperty `json:"properties"` //属性
}
