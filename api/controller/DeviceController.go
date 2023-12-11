package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"smartgw/api/domain"
	"smartgw/api/service"
	"smartgw/lib/io"
)

// DeviceController 设备列表相关
type DeviceController struct {
	service service.DeviceService
}

func NewDeviceController(deviceService service.DeviceService) *DeviceController {
	return &DeviceController{
		service: deviceService,
	}
}

// Add 新增设备列表
func (controller *DeviceController) Add(c *gin.Context) {
	zap.S().Info("DeviceController->Add")
	device := domain.Device{}

	// 解析设备列表
	if c.BindJSON(&device) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备列表失败！", "data": nil})
		return
	}

	// 新增设备列表
	if err := controller.service.Add(&device); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增设备列表失败：" + err.Error(), "data": nil})
		return
	}

	// 新增设备列表成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增设备列表成功！", "data": device})
}

func (controller *DeviceController) Update(c *gin.Context) {
	zap.S().Info("DeviceController->Update")
	device := domain.Device{}
	// 解析设备列表数据
	if c.BindJSON(&device) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备列表失败！", "data": nil})
		return
	}

	// 修改设备列表失败
	if controller.service.Update(&device) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改设备列表失败！", "data": nil})
		return
	}

	// 修改设备列表成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改设备列表成功！", "data": device})
}

func (controller *DeviceController) Delete(c *gin.Context) {
	zap.S().Info("DeviceController->Delete")

	name := c.Param("name")
	if controller.service.Delete(name) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除设备列表失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除设备列表成功！", "data": nil})
}

func (controller *DeviceController) Find(c *gin.Context) {
	zap.S().Info("DeviceController->Find")
	name := c.Param("name")
	device, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找设备列表失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找设备列表成功！", "data": device})
}

// FindAll 获取所有设备列表
func (controller *DeviceController) FindAll(c *gin.Context) {
	zap.S().Info("DeviceController->FindAll")
	devices, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有设备列表失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有设备列表成功！", "data": devices})
}

// Import 数据导出
func (controller *DeviceController) Import(c *gin.Context) {
	zap.S().Info("DeviceController->Import")

	file, err := c.FormFile("FileName")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "文件上传没有FileName字段！", "data": nil})
		return
	}
	filepath := io.GetCurrentPath() + file.Filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "保存上传文件出错！", "data": nil})
		return
	}

	defer os.Remove(filepath)

	data, err := io.ReadFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "读取上传文件失败！", "data": nil})
		return
	}
	var devices []domain.Device
	if err := json.Unmarshal(data, &devices); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "读取上传文件失败！", "data": nil})
		return
	}

	if clears, err := controller.service.FindAll(); err == nil {
		for _, d := range clears {
			controller.service.Delete(d.Name)
		}
	}

	for _, d := range devices {
		controller.service.Add(&d)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "上传设备数据成功！", "data": devices})
}

// Export 数据导出
func (controller *DeviceController) Export(c *gin.Context) {
	zap.S().Info("DeviceController->Export")

	devices, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备列表失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有设备列表成功！", "data": devices})
}
