package worker

import (
	"encoding/json"
	"smartgw/api/domain"
	"smartgw/api/repository"
	"smartgw/lib/collect/channel"
	"smartgw/lib/collect/collector"
	"smartgw/lib/config"
	"smartgw/lib/logger"
	"smartgw/lib/script"
	"time"
)

type Worker interface {
	Start()
	Stop()
	CommandTask(task any)    // 指派高优先级任务
	CollectTask(task any)    // 指派正常任务
	CollectTaskIsFull() bool //采集通道是否为空
}

var _ Worker = (*worker)(nil)

type worker struct {
	channel.PriorityChannel
	collector.Collector
	script.Runner
	deviceRepository repository.DeviceRepository
	dataChan         chan []byte
	stopChan         chan int
	config           *config.Config
}

func NewWorker(
	domainCollector domain.Collector,
	deviceRepository repository.DeviceRepository,
	config *config.Config,
) Worker {
	result := &worker{
		PriorityChannel:  channel.NewPriorityChannel(),
		Collector:        collector.ConnectorFactory(domainCollector),
		Runner:           script.NewLuaRunner(),
		deviceRepository: deviceRepository,
		dataChan:         make(chan []byte, 1024),
		stopChan:         make(chan int, 1),
		config:           config,
	}

	// RPC 命令优先级高
	result.PriorityChannel.SetPriorWorker(result.CommandExecutor)
	// 采集任务优先级低
	result.PriorityChannel.SetNormalWorker(result.CollectExecutor)

	return result
}

func (w *worker) CommandTask(task any) {
	w.PriorityChannel.DispatchPriorTask(task)
}

func (w *worker) CollectTask(task any) {
	w.PriorityChannel.DispatchNormalTask(task)
}

func (w *worker) CollectTaskIsFull() bool {
	return w.PriorityChannel.NormalTaskIsFull()
}

func (w *worker) Start() {
	// 如果不是每次都打开，那么只在这里打开一次
	if !w.config.OpenEveryTime {
		w.Collector.Open(nil)
	}

	go w.BlockRead()
	w.PriorityChannel.Start()
}

func (w *worker) Stop() {
	w.PriorityChannel.Stop()
	w.stopChan <- 1

	// 如果没事每次都打开，那么只在这里关闭一次
	if !w.config.OpenEveryTime {
		w.Collector.Close()
	}
}

// BlockRead 阻塞读取数据
func (w *worker) BlockRead() {
	data := make([]byte, 1024)
	for {
		select {
		case <-w.stopChan:
			logger.Zap.Info("串口数据读取线程正确退出！")
			return
		default:
			count := w.Collector.Read(data)
			if count > 0 {
				w.dataChan <- data[:count]
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (w *worker) CollectExecutor(task any) {
	device, ok := task.(domain.Device)
	if !ok {
		logger.Zap.Error("数据采集执行器无法转换Device参数")
		return
	}
	logger.Zap.Info("开始采集设备: ", device.Name)

	// 备份 device（通过序列化，反序列化深度复制）
	tempBuff, _ := json.Marshal(device)
	deviceCopy := domain.Device{}
	_ = json.Unmarshal(tempBuff, &deviceCopy)

	// 打开驱动
	w.Runner.Open(device.Type.Driver)
	defer w.Runner.Close()

	// 如果每次都打开，那么在这里打开
	if w.config.OpenEveryTime {
		w.Collector.Open(&device)
		defer w.Collector.Close()
	}

	// 数据读取，超过30次，自动退出
	for step := 0; step < 30; step++ {
		data, result, continued := w.Runner.GenerateGetRealVariables(device.Address, step)
		if result {
			// 数据发送
			logger.Zap.Infof("向设备[%s]发送数据[%d:%X]", device.Name, len(data), data)
			w.Collector.Write(data)

			rxBuf := make([]byte, 1024)
			rxBufCnt := 0
			rxTotalBuf := make([]byte, 0)
			rxTotalBufCnt := 0

			reader := func() bool {
				for {
					select {
					case rxBuf = <-w.dataChan:
						rxBufCnt = len(rxBuf)
						if rxBufCnt > 0 {
							rxTotalBufCnt += rxBufCnt
							rxTotalBuf = append(rxTotalBuf, rxBuf[:rxBufCnt]...)
							rxBufCnt = 0
							rxBuf = rxBuf[0:0]
						}

						if rxTotalBufCnt > 0 {
							logger.Zap.Infof("从设备[%s]接收数据[%d:%X]", device.Name, rxTotalBufCnt, rxTotalBuf)
						}

						var tempVariables []domain.DeviceProperty // = make([]domain.DeviceProperty, 0)
						if rxTotalBufCnt > 0 && w.Runner.AnalysisRx(device.Address, device.Type.Properties, rxTotalBuf, rxTotalBufCnt, &tempVariables) {
							logger.Zap.Infof("从设备[%s]接收数据后，正确解析数据", device.Name)
							device.CollectTime = time.Now()
							device.Online = true
							device.CollectTotal += 1
							device.CollectSuccess += 1
							w.deviceRepository.Save(&device)
							led.SetRunningNormal()
							return false
						}
					case <-time.After(time.Duration(w.Collector.GetTimeout()) * time.Millisecond):
						if rxTotalBufCnt > 0 {
							logger.Zap.Infof("[超时]从设备[%s]接收数据[%d:%X]", device.Name, rxTotalBufCnt, rxTotalBuf)
						}

						var tempVariables []domain.DeviceProperty // = make([]domain.DeviceProperty, 0)
						if rxTotalBufCnt > 0 && w.Runner.AnalysisRx(device.Address, device.Type.Properties, rxTotalBuf, rxTotalBufCnt, &tempVariables) {
							logger.Zap.Infof("从设备[%s]接收数据后，超时，但正确解析数据", device.Name)
							device.CollectTime = time.Now()
							device.Online = true
							device.CollectTotal += 1
							device.CollectSuccess += 1
							w.deviceRepository.Save(&device)
							led.SetRunningNormal()
							return false
						}

						// 此处超时，可以设置设备离线
						deviceCopy.CollectTime = time.Now()
						deviceCopy.Online = false
						deviceCopy.CollectTotal += 1
						w.deviceRepository.Save(&deviceCopy)
						led.SetRunningSerialErr()
						return true
					}
				}
			}
			if timeout := reader(); timeout {
				// 如果一次读取超时，那么我们认为读取该设备的其他属性也会超时，所以就退出了
				logger.Zap.Error("超时退出")
				break
			}
		}
		if !continued {
			break
		}
		time.Sleep(time.Duration(w.Collector.GetInterval()) * time.Millisecond)
	}

	logger.Zap.Info("结束采集设备:", device.Name)
}

func (w *worker) CommandExecutor(task any) {
	commandRequest, ok := task.(CommandRequest)
	if !ok {
		logger.Zap.Error("RPC执行器无法转换CommandRequest参数")
		return
	}
	// 串口测试
	if commandRequest.Method == "test" {
		w.CommandTest(commandRequest)
		return
	}
	logger.Zap.Info("开始执行RPC命令:", commandRequest.Method)

	requestParam := commandRequest.Params[0]

	device, err := w.deviceRepository.Find(requestParam.ClientID)

	if err != nil {
		return
	}

	// 打开设备驱动
	w.Runner.Open(device.Type.Driver)
	defer w.Runner.Close()

	// 如果每次都打开，那么在这里打开
	if w.config.OpenEveryTime {
		w.Collector.Open(&device)
		defer w.Collector.Close()
	}

	// 只有最后一条有返回结果
	responseParam := ResponseParam{}
	responseParam.ClientID = requestParam.ClientID
	responseParam.CmdName = requestParam.CmdName
	responseParam.CmdStatus = 1
	responseParam.CmdResult = nil
	// 至少有一条参数（命令）
	for _, requestParam = range commandRequest.Params {
		responseParam.CmdName = requestParam.CmdName
		device, err = w.deviceRepository.Find(requestParam.ClientID)

		if err != nil {
			break
		}
		// 数据读取，超过30次，自动退出
		for step := 0; step < 30; step++ {
			cmdParams, _ := json.Marshal(&(requestParam.CmdParams))
			data, result, continued := w.Runner.DeviceCustomCmd(device.Address, requestParam.CmdName, string(cmdParams), step)
			if result {
				// 数据发送
				logger.Zap.Infof("RPC向设备[%s]发送数据[%d:%X]", device.Name, len(data), data)
				w.Collector.Write(data)

				rxBuf := make([]byte, 1024)
				rxBufCnt := 0
				rxTotalBuf := make([]byte, 0)
				rxTotalBufCnt := 0

				reader := func() bool {
					for {
						select {
						case rxBuf = <-w.dataChan:
							rxBufCnt = len(rxBuf)
							if rxBufCnt > 0 {
								rxTotalBufCnt += rxBufCnt
								rxTotalBuf = append(rxTotalBuf, rxBuf[:rxBufCnt]...)
								rxBufCnt = 0
								rxBuf = rxBuf[0:0]
							}
							if rxTotalBufCnt > 0 {
								logger.Zap.Infof("从设备[%s]接收数据[%d:%X]", device.Name, rxTotalBufCnt, rxTotalBuf)
							}
							var tempVariables []domain.DeviceProperty //= make([]domain.DeviceProperty, 0)
							if rxTotalBufCnt > 0 && w.Runner.AnalysisRx(device.Address, device.Type.Properties, rxTotalBuf, rxTotalBufCnt, &tempVariables) {
								responseParam.CmdStatus = 0
								tempMap := make(map[string]interface{})
								for _, property := range tempVariables {
									tempMap[property.Name] = property.Value
								}
								responseParam.CmdResult = tempMap
								return false
							}
						case <-time.After(time.Duration(w.Collector.GetTimeout()) * time.Millisecond):
							if rxTotalBufCnt > 0 {
								logger.Zap.Infof("从设备[%s]接收数据[%d:%X]", device.Name, rxTotalBufCnt, rxTotalBuf)
							}
							var tempVariables []domain.DeviceProperty //= make([]domain.DeviceProperty, 0)
							if rxTotalBufCnt > 0 && w.Runner.AnalysisRx(device.Address, device.Type.Properties, rxTotalBuf, rxTotalBufCnt, &tempVariables) {
								responseParam.CmdStatus = 0
								tempMap := make(map[string]interface{})
								for _, property := range tempVariables {
									tempMap[property.Name] = property.Value
								}
								responseParam.CmdResult = tempMap
								return false
							}
							responseParam.CmdStatus = 1
							responseParam.CmdResult = nil
							return true
						}
					}
				}
				if timeout := reader(); timeout {
					// 如果一次读取超时，那么我们认为读取该设备的其他属性也会超时，所以就退出了
					logger.Zap.Error("超时退出")
					break
				}
			}
			if !continued {
				break
			}
			time.Sleep(time.Duration(w.Collector.GetInterval()) * time.Millisecond)
		}
	}

	commandRequest.ResponseParamChan <- responseParam

	logger.Zap.Info("结束执行RPC命令:", commandRequest.Method)
}

func (w *worker) CommandTest(commandRequest CommandRequest) {
	// 如果每次都打开，那么在这里打开
	if w.config.OpenEveryTime {
		logger.Zap.Error("每次都打开设备，不支持调试!")
		return
	}

	if commandRequest.Method != "test" {
		return
	}

	requestParam := commandRequest.Params[0]

	logger.Zap.Info("开始执行测试命令:", requestParam.CmdName)

	responseParam := ResponseParam{}

	responseParam.CmdStatus = 0
	responseParam.CmdResult = nil

	// 发送指令
	data, ok := requestParam.CmdParams["param"].([]byte)
	if !ok {
		logger.Zap.Error("串口测试无法获取参数")
		return
	}
	w.Collector.Write(data)
	logger.Zap.Infof("发送【测试】数据[%d:%X]", len(data), data)
	rxBuf := make([]byte, 1024)
	rxBufCnt := 0
	rxTotalBuf := make([]byte, 0)
	rxTotalBufCnt := 0

	// 接收指令
	func() {
		for {
			select {
			case rxBuf = <-w.dataChan:
				rxBufCnt = len(rxBuf)
				if rxBufCnt > 0 {
					rxTotalBufCnt += rxBufCnt
					rxTotalBuf = append(rxTotalBuf, rxBuf[:rxBufCnt]...)
					rxBufCnt = 0
					rxBuf = rxBuf[0:0]
				}
			case <-time.After(time.Duration(w.Collector.GetTimeout()) * time.Millisecond):
				responseParam.CmdResult = rxTotalBuf[:rxTotalBufCnt]
				if rxTotalBufCnt == 0 {
					responseParam.CmdStatus = 1
				} else {
					responseParam.CmdStatus = 0
					if rxTotalBufCnt > 0 {
						logger.Zap.Infof("接收【测试】数据[%d:%X]", rxTotalBufCnt, rxTotalBuf)
					}
				}
				return
			}
		}
	}()

	commandRequest.ResponseParamChan <- responseParam
	logger.Zap.Infof("结束执行测试命令:", requestParam.CmdName)
}
