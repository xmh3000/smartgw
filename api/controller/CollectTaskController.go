package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
)

// CollectTaskController 采集任务
type CollectTaskController struct {
	service service.CollectTaskService
}

func NewCollectTaskController(collectTaskService service.CollectTaskService) *CollectTaskController {
	return &CollectTaskController{
		service: collectTaskService,
	}
}

// Add 新增采集任务
func (controller *CollectTaskController) Add(c *gin.Context) {
	zap.S().Info("CollectTaskController->Add")
	collectTask := domain.CollectTask{}

	// 解析采集任务
	if c.BindJSON(&collectTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析采集任务数据失败！", "data": nil})
		return
	}

	// 新增采集任务失败
	if err := controller.service.Add(&collectTask); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增采集任务失败：" + err.Error(), "data": nil})
		return
	}

	// 新增采集任务成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增采集任务成功！", "data": collectTask})
}

func (controller *CollectTaskController) Update(c *gin.Context) {
	zap.S().Info("CollectTaskController->Update")
	collectTask := domain.CollectTask{}
	// 解析采集任务数据
	if c.BindJSON(&collectTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析采集任务数据失败！", "data": nil})
		return
	}

	// 修改采集任务失败
	if controller.service.Update(&collectTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改采集任务失败！", "data": nil})
		return
	}

	// 修改采集任务成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改采集任务成功！", "data": collectTask})
}

func (controller *CollectTaskController) Delete(c *gin.Context) {
	zap.S().Info("CollectTaskController->Delete")
	name := c.Param("name")
	if controller.service.Delete(name) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除采集任务失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除采集任务成功！", "data": nil})
}

func (controller *CollectTaskController) Find(c *gin.Context) {
	zap.S().Info("CollectTaskControllerController->Find")
	name := c.Param("name")
	collectTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找采集任务失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找采集任务成功！", "data": collectTask})
}

// FindAll 获取所有采集任务数据
func (controller *CollectTaskController) FindAll(c *gin.Context) {
	zap.S().Info("CollectTaskController->FindAll")
	collectTasks, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有采集任务数据失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有采集任务数据成功！", "data": collectTasks})
}

func (controller *CollectTaskController) Start(c *gin.Context) {
	zap.S().Info("CollectTaskController->Start")

	name := c.Param("name")
	reportTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "找不到相关服务:" + name, "data": nil})
		return
	}

	reportTask.Status = 1

	err = controller.service.Update(&reportTask)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改任务状态失败:" + name, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "启动任务成功！", "data": reportTask})
}

func (controller *CollectTaskController) Stop(c *gin.Context) {
	zap.S().Info("CollectTaskController->Stop")

	name := c.Param("name")
	reportTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "找不到相关服务:" + name, "data": nil})
		return
	}

	reportTask.Status = 0

	err = controller.service.Update(&reportTask)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改任务状态失败:" + name, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "停止任务成功！", "data": reportTask})
}
