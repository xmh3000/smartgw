package domain

// Serial 本地串口，主要指两个485串口
type Serial struct {
	Name       string `json:"name"`       //串口名称
	DeviceName string `json:"deviceName"` //设备名称
	BaudRate   int    `json:"baudRate"`   //波特率
	DataBit    int    `json:"dataBit"`    //数据位
	StopBit    string `json:"stopBit"`    //停止位
	Check      string `json:"check"`      //检验
}

// TcpClient 南向网口，主要指接入网桥的设备，即485转tcp协议
type TcpClient struct {
	Name string `json:"name"` //名称
	Ip   string `json:"ip"`   //Ip地址
	Port int    `json:"port"` //端口号
}

// Collector 采集接口，目前包含本地串口、TcpClient，未来可能包含开关量
type Collector struct {
	Name      string    `json:"name"`      //接口名称
	Type      string    `json:"type"`      //接口类型Serial或TcpClient
	Serial    Serial    `json:"serial"`    //串口
	TcpClient TcpClient `json:"tcpClient"` //Tcp客户端
	Timeout   int       `json:"timeout"`   //超时
	Interval  int       `json:"interval"`  //间隔
}
