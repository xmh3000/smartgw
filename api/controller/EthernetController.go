package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
)

type EthernetController struct {
	service service.EthernetService
}

func NewEthernetController(ethernetService service.EthernetService) *EthernetController {
	return &EthernetController{
		service: ethernetService,
	}
}

func (controller *EthernetController) Add(c *gin.Context) {
	zap.S().Info("EthernetController->Add")
	ethernet := domain.Ethernet{}

	if c.BindJSON(&ethernet) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析网口失败", "data": nil})
		return
	}

	if err := controller.service.Add(&ethernet); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增网口失败", "data": ethernet})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增网口成功！", "data": ethernet})
}

func (controller *EthernetController) Update(c *gin.Context) {
	zap.S().Info("EthernetController->Update")
	ethernet := domain.Ethernet{}

	if c.BindJSON(&ethernet) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析网口失败！", "data": nil})
		return
	}

	if controller.service.Update(&ethernet) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改网口失败！", "data": ethernet})
		return
	}

	// 修改设备列表成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改网口成功！", "data": ethernet})
}

func (controller *EthernetController) Delete(c *gin.Context) {
	zap.S().Info("EthernetController->Delete")

	name := c.Param("name")
	if controller.service.Delete(name) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除网口失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除网口成功！", "data": nil})
}

func (controller *EthernetController) Find(c *gin.Context) {
	zap.S().Info("EthernetController->Find")
	name := c.Param("name")
	ethernet, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找网口失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找网口成功！", "data": ethernet})
}

// FindAll 获取所有设备列表
func (controller *EthernetController) FindAll(c *gin.Context) {
	zap.S().Info("EthernetController->FindAll")
	ethernets, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有网口失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有网口成功！", "data": ethernets})
}
