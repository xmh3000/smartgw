package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
)

// CollectorController 采集接口控制器
type CollectorController struct {
	service service.CollectorService
}

func NewCollectorController(collectorService service.CollectorService) *CollectorController {
	return &CollectorController{
		service: collectorService,
	}
}

// Add 新增采集接口
func (controller *CollectorController) Add(c *gin.Context) {
	zap.S().Info("CollectorController->Add")
	collector := domain.Collector{}

	// 解析采集接口数据
	if c.BindJSON(&collector) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析采集接口数据失败！", "data": nil})
		return
	}

	// 新增采集接口失败
	if err := controller.service.Add(&collector); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增采集接口失败：" + err.Error(), "data": nil})
		return
	}

	// 新增采集接口成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增采集接口成功！", "data": collector})
}

func (controller *CollectorController) Update(c *gin.Context) {
	zap.S().Info("CollectorController->Update")
	collector := domain.Collector{}
	// 解析采集接口数据
	if c.BindJSON(&collector) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析采集接口数据失败！", "data": nil})
		return
	}

	// 修改采集接口失败
	if controller.service.Update(&collector) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改采集接口失败！", "data": nil})
		return
	}

	// 修改采集接口成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改采集接口成功！", "data": collector})
}

func (controller *CollectorController) Delete(c *gin.Context) {
	zap.S().Info("CollectorController->Delete")

	name := c.Param("name")

	if err := controller.service.Delete(name); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除采集接口失败：" + err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除采集接口成功！", "data": nil})
}

func (controller *CollectorController) Find(c *gin.Context) {
	zap.S().Info("CollectorController->Find")
	name := c.Param("name")

	collector, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找采集接口失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找采集接口成功！", "data": collector})
}

// FindAll 获取所有采集接口数据
func (controller *CollectorController) FindAll(c *gin.Context) {
	zap.S().Info("UserController->FindAll")
	collectors, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有采集接口数据失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有采集接口数据成功！", "data": collectors})
}
