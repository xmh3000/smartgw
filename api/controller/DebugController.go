package controller

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"smartgw/api/repository"
	"smartgw/lib/collect"
	"smartgw/lib/collect/worker"
	"smartgw/lib/config"
	"smartgw/lib/io"
	"strings"
	"time"
)

type DebugController struct {
	collectorServer      *collect.CollectorServer
	deviceRepository     repository.DeviceRepository
	reportTaskRepository repository.ReportTaskRepository
	config               *config.Config
}

func NewDebugController(
	collectorServer *collect.CollectorServer,
	deviceRepository repository.DeviceRepository,
	reportTaskRepository repository.ReportTaskRepository,
	config *config.Config) *DebugController {
	return &DebugController{
		collectorServer:      collectorServer,
		deviceRepository:     deviceRepository,
		reportTaskRepository: reportTaskRepository,
		config:               config,
	}
}

func (controller *DebugController) Test(c *gin.Context) {
	//controller.c = c

	zap.S().Info("DebugController->test()")

	debugRequest := struct {
		CollectorName string `json:"collectorName"`
		Data          string `json:"data"`
	}{}

	if c.BindJSON(&debugRequest) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析采集接口数据失败！", "data": nil})
		return
	}

	debugRequest.Data = strings.ReplaceAll(debugRequest.Data, " ", "")
	if len(debugRequest.Data)%2 != 0 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "16进制数据格式不正确，必须是偶数个数字！", "data": nil})
		return
	}

	commandRequest := worker.CommandRequest{
		Method: "test",
		Params: []worker.RequestParam{
			{
				ClientID:  debugRequest.CollectorName,
				CmdName:   "test",
				CmdParams: make(map[string]interface{}),
			},
		},
	}
	reqData, _ := hex.DecodeString(debugRequest.Data)
	commandRequest.Params[0].CmdParams["param"] = reqData
	//commandRequest.Callback = controller.Callback
	commandRequest.ResponseParamChan = make(chan worker.ResponseParam, 1)
	if worker, ok := controller.collectorServer.FindByCollectorName(debugRequest.CollectorName); ok {
		worker.CommandTask(commandRequest)

		response := <-commandRequest.ResponseParamChan
		if response.CmdStatus == 0 {
			data := response.CmdResult.([]byte)
			c.JSON(http.StatusOK, gin.H{"code": response.CmdStatus, "message": "正确应答", "data": hex.EncodeToString(data)})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": response.CmdStatus, "message": "超时", "data": nil})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "超时", "data": nil})
	}
}

// 升级相关
func (controller *DebugController) Upgrade(c *gin.Context) {
	zap.S().Debug("准备升级")
	file, err := c.FormFile("FileName")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "文件上传没有FileName字段！", "data": nil})
		return
	}
	// 将压缩包移动到当前运行目录 filepath: C:\Users\w\AppData\Local\Temp\GoLand\up.zip
	filepath := io.GetCurrentPath() + file.Filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "保存上传文件出错！", "data": nil})
		return
	}

	err = io.Unzip(filepath, io.GetCurrentPath())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解压缩文件出错！", "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "升级成功，系统重启！", "data": nil})
	os.Remove(filepath)
	io.OsCommand("chmod +x " + io.GetCurrentPath() + "smartgw")
	io.SystemReboot()
}

func (controller *DebugController) SystemStatus(c *gin.Context) {
	io.SystemState.Name = controller.config.ClientID
	if controller.config.OpenEveryTime {
		io.SystemState.OpenEveryTime = "支持不同波特率"
	} else {
		io.SystemState.OpenEveryTime = "不支持不同波特率"
	}
	io.GetMemState()
	io.GetDiskState()
	io.GetRuntime()
	//设备在线率计算
	deviceTotalCnt := 0
	deviceOnlineCnt := 0
	devices, _ := controller.deviceRepository.FindAll()
	zap.S().Debug(len(devices))
	for _, device := range devices {
		zap.S().Debug(device)
		deviceTotalCnt = deviceTotalCnt + 1
		if device.Online == true {
			deviceOnlineCnt = deviceOnlineCnt + 1
		}
	}

	if deviceOnlineCnt == 0 {
		io.SystemState.DeviceTotal = "0"
		io.SystemState.DeviceOnline = "0"
	} else {
		io.SystemState.DeviceTotal = fmt.Sprintf("%d", deviceTotalCnt)
		io.SystemState.DeviceOnline = fmt.Sprintf("%5.2f", float32(deviceOnlineCnt*100.0/deviceTotalCnt))
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": io.SystemState})
}

func (controller *DebugController) SystemReboot(c *gin.Context) {
	io.SystemReboot()
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": ""})
}

func (controller *DebugController) SystemNtp(c *gin.Context) {
	io.ExecuteNtp(controller.config)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": timeStr})
}

//
//func (controller *DebugController) WatchdogStart(c *gin.Context) {
//	if controller.watchdog.Running() != true {
//		controller.watchdog.Start()
//	}
//	if controller.watchdog.Running() == true {
//		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "开启看门狗成功", "data": ""})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "开启看门狗失败", "data": ""})
//	}
//}
//
//func (controller *DebugController) WatchdogStop(c *gin.Context) {
//	if controller.watchdog.Running() != false {
//		controller.watchdog.Start()
//	}
//	if controller.watchdog.Running() == false {
//		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "关闭看门狗成功", "data": ""})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "关闭看门狗失败", "data": ""})
//	}
//}
