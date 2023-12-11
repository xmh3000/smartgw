package mqtt

import (
	"go.uber.org/zap"
	"smartgw/api/domain"
	"smartgw/lib/io"
	"time"
)

// Device 设备信息，只用于数据上报，下发。避免与api/domain/Device命名冲突
type Device struct {
	CollectName string `json:"collectName"` // 采集接口名称
	Name        string `json:"name"`        //
	Addr        string `json:"addr"`
	Type        string `json:"type"`
}

type Device2 struct {
	CollectName string        `json:"collectName"` // 采集接口名称
	Name        string        `json:"name"`        //
	Addr        string        `json:"addr"`
	Type        string        `json:"type"`
	Alone       bool          `json:"alone"`  //独立开关
	Serial      domain.Serial `json:"serial"` //通讯参数
}

// 处理网关操作 网关重启(reboot)、远程升级(upgrade)
// 获取表具设备(getdevice)、下发表具设备(setdevice)、PING(ping)
const (
	REBOOT     = "Reboot"
	UPGRADE    = "Upgrade"
	GETDEVICE  = "GetDevices"
	SETDEVICE  = "SetDevices"
	GETDEVICE2 = "GetDevices2"
	SETDEVICE2 = "SetDevices2"
	PING       = "Ping"
)

type commandExecutor func(params map[string]interface{}, client Client) (int, interface{})

var (
	GatewayCommand map[string]commandExecutor = initGatewayCommand()
)

func initGatewayCommand() map[string]commandExecutor {
	result := make(map[string]commandExecutor)

	result[REBOOT] = reboot
	result[UPGRADE] = upgrade
	result[GETDEVICE] = getDevice
	result[SETDEVICE] = setDevice
	result[GETDEVICE2] = getDevice2
	result[SETDEVICE2] = setDevice2
	result[PING] = ping

	return result
}

func reboot(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令reboot")

	io.SystemReboot()

	return 0, nil
}

func upgrade(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令upgrade")

	sourceFile, ok := params["url"].(string)
	if !ok {
		zap.S().Error("升级失败，缺少url参数")
		return 1, nil
	}

	fileName := "smartgw.zip"
	if len(sourceFile) > 0 {
		io.Upgrade(sourceFile, io.GetCurrentPath()+fileName)
	}

	return 0, nil
}

func getDevice(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令GetDevices")

	var result []Device

	devices, err := client.deviceRepository.FindAll()
	if err != nil {
		zap.S().Error("RPC获取设备失败")
		return 1, result
	}

	for _, d := range devices {
		device := Device{
			CollectName: d.Collector.Name,
			Name:        d.Name,
			Addr:        d.Address,
			Type:        d.Type.Name,
		}
		result = append(result, device)
	}

	return 0, result
}

func getDevice2(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令GetDevices2")

	var result []Device2

	devices, err := client.deviceRepository.FindAll()
	if err != nil {
		zap.S().Error("RPC获取设备失败")
		return 1, result
	}

	for _, d := range devices {
		device := Device2{
			CollectName: d.Collector.Name,
			Name:        d.Name,
			Addr:        d.Address,
			Type:        d.Type.Name,
			Alone:       d.Alone,
			Serial:      d.Serial,
		}
		result = append(result, device)
	}

	return 0, result
}

func setDevice(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令SetDevices")

	// 1. 获取设备列表
	devices := convertDevices(params)

	if devices == nil || len(devices) == 0 {
		return 0, nil
	}

	// 2. 删除设备信息
	if clears, err := client.deviceRepository.FindAll(); err == nil {
		for _, d := range clears {
			// 删除
			client.deviceRepository.Delete(d.Name)
		}
	}

	// 3. 添加设备信息
	for _, d := range devices {
		deviceType, err := client.deviceTypeRepository.Find(d.Type)
		if err != nil {
			zap.S().Debug("设备类型[" + d.Type + "]不存在！")
			continue
		}
		collector, err := client.collectorRepository.Find(d.CollectName)

		if err != nil {
			zap.S().Debug("采集接口[" + d.CollectName + "]不存在！")
			continue
		}

		device := &domain.Device{
			Name:           d.Name,
			Type:           deviceType,
			Address:        d.Addr,
			Collector:      collector,
			Alone:          false,
			Serial:         collector.Serial,
			Online:         false,
			CollectTime:    time.Time{},
			CollectTotal:   0,
			CollectSuccess: 0,
			ReportTime:     time.Time{},
			ReportTotal:    0,
			ReportSuccess:  0,
		}
		client.deviceRepository.Save(device)
	}

	return 0, nil
}

// ------------------------------------------------------------
func convertDevices(params map[string]interface{}) (result []Device) {
	result = []Device{}

	if devices, ok := params["devices"]; ok {
		if devs, ok := devices.([]interface{}); ok {
			for _, device := range devs {
				if dev, ok := device.(map[string]interface{}); ok {
					d := Device{}

					if collectName, ok := dev["collectName"]; ok {
						d.CollectName = convertString(collectName)
					}

					if name, ok := dev["name"]; ok {
						d.Name = convertString(name)
					}

					if addr, ok := dev["addr"]; ok {
						d.Addr = convertString(addr)
					}

					if type2, ok := dev["type"]; ok {
						d.Type = convertString(type2)
					}

					result = append(result, d)
				}
			}
		}
	}
	return result
}

func setDevice2(params map[string]interface{}, client Client) (int, interface{}) {
	zap.S().Info("执行RPC命令SetDevices2")

	// 1. 获取设备列表
	devices := convertDevices2(params)

	if devices == nil || len(devices) == 0 {
		return 1, nil
	}

	// 2. 删除设备信息
	if clears, err := client.deviceRepository.FindAll(); err == nil {
		for _, d := range clears {
			// 删除
			client.deviceRepository.Delete(d.Name)
		}
	}

	// 3. 添加设备信息
	for _, d := range devices {
		deviceType, err := client.deviceTypeRepository.Find(d.Type)
		if err != nil {
			zap.S().Debug("设备类型[" + d.Type + "]不存在！")
			continue
		}
		collector, err := client.collectorRepository.Find(d.CollectName)

		if err != nil {
			zap.S().Debug("采集接口[" + d.CollectName + "]不存在！")
			continue
		}

		device := &domain.Device{
			Name:           d.Name,
			Type:           deviceType,
			Address:        d.Addr,
			Collector:      collector,
			Alone:          d.Alone,
			Serial:         d.Serial,
			Online:         false,
			CollectTime:    time.Time{},
			CollectTotal:   0,
			CollectSuccess: 0,
			ReportTime:     time.Time{},
			ReportTotal:    0,
			ReportSuccess:  0,
		}

		client.deviceRepository.Save(device)
	}

	return 0, nil
}

// ------------------------------------------------------------
func convertDevices2(params map[string]interface{}) (result []Device2) {
	result = []Device2{}

	if devices, ok := params["devices"]; ok {
		if devs, ok := devices.([]interface{}); ok {
			for _, device := range devs {
				if dev, ok := device.(map[string]interface{}); ok {
					d := Device2{}

					if collectName, ok := dev["collectName"]; ok {
						d.CollectName = convertString(collectName)
					}

					if name, ok := dev["name"]; ok {
						d.Name = convertString(name)
					}

					if addr, ok := dev["addr"]; ok {
						d.Addr = convertString(addr)
					}

					if type2, ok := dev["type"]; ok {
						d.Type = convertString(type2)
					}

					if alone, ok := dev["alone"]; ok {
						d.Alone = convertBoolean(alone)
					}

					if d.Alone {
						if serial, ok := dev["serial"]; ok {
							d.Serial = *convertSerial(serial)
						}
					}

					result = append(result, d)
				}
			}
		}
	}
	return
}

func convertSerial(value interface{}) (result *domain.Serial) {
	result = &domain.Serial{}

	if serial, ok := value.(map[string]interface{}); ok {
		if name, ok := serial["name"]; ok {
			result.Name = convertString(name)
		}

		if deviceName, ok := serial["deviceName"]; ok {
			result.DeviceName = convertString(deviceName)
		}

		if baudRate, ok := serial["baudRate"]; ok {
			result.BaudRate = (int)(convertFloat64(baudRate))
		}
		if dataBit, ok := serial["dataBit"]; ok {
			result.DataBit = (int)(convertFloat64(dataBit))
		}

		if stopBit, ok := serial["stopBit"]; ok {
			result.StopBit = convertString(stopBit)
		}

		if check, ok := serial["check"]; ok {
			result.Check = convertString(check)
		}
	}

	return
}

func convertString(value interface{}) (result string) {
	result = ""

	if temp, ok := value.(string); ok {
		result = temp
	}

	return
}

func convertFloat64(value interface{}) (result float64) {
	result = 0

	if temp, ok := value.(float64); ok {
		result = temp
	}

	return
}

func convertBoolean(value interface{}) (result bool) {
	result = false

	if temp, ok := value.(bool); ok {
		result = temp
	}

	return
}

func ping(params map[string]interface{}, client Client) (int, interface{}) {
	return 0, nil
}
