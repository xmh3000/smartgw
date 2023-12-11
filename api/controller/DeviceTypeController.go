package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path"
	"smartgw/api/domain"
	"smartgw/api/service"
	"smartgw/lib/io"
	"strconv"
	"strings"
)

// DeviceTypeController 对设备类型的相关操作
type DeviceTypeController struct {
	service service.DeviceTypeService
}

func NewDeviceTypeController(deviceTypeService service.DeviceTypeService) *DeviceTypeController {
	return &DeviceTypeController{
		service: deviceTypeService,
	}
}

// Add 新增设备类型
func (controller *DeviceTypeController) Add(c *gin.Context) {
	zap.S().Info("DeviceTypeController->Add")
	devieType := domain.DeviceType{}

	// 解析设备类型数据
	if c.BindJSON(&devieType) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备类型数据失败！", "data": nil})
		return
	}

	// 新增设备类型失败
	if err := controller.service.Add(&devieType); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增设备类型失败：" + err.Error(), "data": nil})
		return
	}

	// 新增设备类型成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增设备类型成功！", "data": devieType})
}

func (controller *DeviceTypeController) Update(c *gin.Context) {
	zap.S().Info("DeviceTypeController->Update")
	deviceType := domain.DeviceType{}
	// 解析用户数据
	if c.BindJSON(&deviceType) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备类型数据失败！", "data": nil})
		return
	}

	// 修改设备类型失败
	if controller.service.Update(&deviceType) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改设备类型失败！", "data": nil})
		return
	}

	// 修改设备类型成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改设备类型成功！", "data": deviceType})
}

func (controller *DeviceTypeController) Delete(c *gin.Context) {
	zap.S().Info("DeviceTypeController->Delete")

	name := c.Param("name")
	if err := controller.service.Delete(name); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除设备类型失败：" + err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除设备类型成功！", "data": nil})
}

func (controller *DeviceTypeController) Find(c *gin.Context) {
	zap.S().Info("DeviceTypeController->Find")
	name := c.Param("name")
	deviceType, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找设备类型失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找设备类型成功！", "data": deviceType})
}

// FindAll 获取所有设备类型数据
func (controller *DeviceTypeController) FindAll(c *gin.Context) {
	zap.S().Info("DeviceTypeController->FindAll")
	deviceTypes, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有设备类型数据失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有设备类型数据成功！", "data": deviceTypes})
}

// AddProperties 增加设备属性
func (controller *DeviceTypeController) AddProperties(c *gin.Context) {
	zap.S().Info("DeviceTypeController->AddProperties")
	name := c.Param("name")
	deviceProperty := domain.DeviceProperty{}

	if c.BindJSON(&deviceProperty) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备属性失败！", "data": nil})
		return
	}

	// 新增设备属性
	if controller.service.AddProperties(name, &deviceProperty) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增设备属性失败！", "data": nil})
		return
	}

	// 新增设备属性成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增设备属性成功！", "data": deviceProperty})
}

// UpdateProperties 修改设备属性
func (controller *DeviceTypeController) UpdateProperties(c *gin.Context) {
	zap.S().Info("DeviceTypeController->UpdateProperties")
	name := c.Param("name")

	//获取设备属性id
	propertyIdParam := c.Param("propertyid")
	propertyId, err := strconv.ParseInt(propertyIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备类型id失败！", "data": nil})
		return
	}

	deviceProperty := domain.DeviceProperty{}
	if c.BindJSON(&deviceProperty) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析设备属性失败！", "data": nil})
		return
	}

	if controller.service.UpdateProperties(name, int(propertyId), &deviceProperty) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改设备属性失败！", "data": nil})
		return
	}

	// 修改设备属性成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改设备属性成功！", "data": deviceProperty})
}

// DeleteProperties 删除设备属性
func (controller *DeviceTypeController) DeleteProperties(c *gin.Context) {
	zap.S().Info("DeviceTypeController->DeleteProperties")
	name := c.Param("name")

	//获取设备属性id
	propertyIdParam := c.Param("propertyid")
	propertyId, err := strconv.ParseInt(propertyIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备类型id失败！", "data": nil})
		return
	}

	if controller.service.DeleteProperties(name, int(propertyId)) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除设备属性失败！", "data": nil})
		return
	}

	// 删除设备属性成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除设备属性成功！", "data": nil})
}

// FindProperty 查找某一设备某一个属性
func (controller *DeviceTypeController) FindProperty(c *gin.Context) {
	zap.S().Info("DeviceTypeController->FindProperty")
	name := c.Param("name")

	//获取设备属性id
	propertyIdParam := c.Param("propertyid")
	propertyId, err := strconv.ParseInt(propertyIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备类型id失败！", "data": nil})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备类型id失败！", "data": nil})
		return
	}

	deviceProperty, err := controller.service.FindProperty(name, int(propertyId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找设备属性失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找设备属性成功！", "data": deviceProperty})

}

// FindAllProperties 查找某一个设备的所有属性
func (controller *DeviceTypeController) FindAllProperties(c *gin.Context) {
	zap.S().Info("DeviceTypeController->FindAllProperties")
	name := c.Param("name")

	deviceProperties, err := controller.service.FindAllProperties(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找设备属性失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找设备属性成功！", "data": deviceProperties})

}

func (controller *DeviceTypeController) Upload(c *gin.Context) {
	file, err := c.FormFile("FileName")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "文件上传没有FileName字段！", "data": nil})
		return
	}

	// 确保文件夹plugin存在
	io.SureExists("plugin")
	dir := io.GetCurrentPath() + "plugin/"
	filename := dir + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "保存上传文件出错！", "data": nil})
		return
	}
	defer os.Remove(filename)

	err = io.Unzip(filename, dir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解压缩文件出错！", "data": nil})
		return
	}

	name := c.Param("name")
	deviceType, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取设备类型失败！", "data": nil})
		return
	}
	filenameWithSuffix := path.Base(filename)
	fileType := path.Ext(filename)
	plugin := strings.TrimSuffix(filenameWithSuffix, fileType)
	deviceType.Driver = plugin

	propertiesFile := "plugin/" + plugin + "/properties.json"
	if ok, _ := io.PathExists(propertiesFile); ok {
		if data, err := io.ReadFile(propertiesFile); err == nil {
			properties := []domain.DeviceProperty{}
			if json.Unmarshal(data, &properties) == nil {
				deviceType.Properties = properties
			}
		}
	}

	if err = controller.service.Update(&deviceType); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改设备驱动失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "上传设备驱动成功！", "data": nil})
}
