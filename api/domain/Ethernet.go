package domain

// Ethernet 采集网口，主要用来接网桥
type Ethernet struct {
	Name    string `json:"name"`  // e.g., "en0", "lo0", "eth0.100"
	Index   int    `json:"index"` // positive integer that starts at one, zero is never used
	MTU     int    `json:"MTU"`   // maximum transmission unit
	MAC     string `json:"MAC"`   // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags   string `json:"flags"` // e.g., FlagUp, FlagLoopback, FlagMulticast
	IP      string `json:"IP"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	//--------------------------------------------------------------------------------------------------------
	ConfigEnabled bool   `json:"configEnabled"` // 启用配置
	DHCPEnabled   bool   `json:"DHCPEnabled"`   // 是否手动：true 手动模式; false 自动模式 //     true: 自动模式（DHCP模式），自动获取IP；false：手工模式：手工设置IP
	ConfigIP      string `json:"configIP"`
	ConfigNetmask string `json:"configNetmask"`
	ConfigGateway string `json:"configGateway"`
}
