package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
)

// ReportTaskController 上报服务相关控制器
type ReportTaskController struct {
	service service.ReportTaskService
}

func NewReportTaskController(reportTaskService service.ReportTaskService) *ReportTaskController {
	return &ReportTaskController{
		service: reportTaskService,
	}
}

// Add 新增上报服务
func (controller *ReportTaskController) Add(c *gin.Context) {
	zap.S().Info("ReportTaskController->Add")
	reportTask := domain.ReportTask{}

	// 解析上报服务数据
	if c.BindJSON(&reportTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析上报服务数据失败！", "data": nil})
		return
	}

	// 新增上报服务失败
	if controller.service.Add(&reportTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增上报服务失败！", "data": nil})
		return
	}

	// 新增上报服务成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增上报服务成功！", "data": reportTask})
}

func (controller *ReportTaskController) Update(c *gin.Context) {
	zap.S().Info("ReportTaskController->Update")
	reportTask := domain.ReportTask{}
	// 解析上报服务数据
	if c.BindJSON(&reportTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析上报服务失败！", "data": nil})
		return
	}

	// 修改上报服务失败
	if controller.service.Update(&reportTask) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改上报服务失败！", "data": nil})
		return
	}

	// 修改上报服务成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改上报服务成功！", "data": reportTask})
}

func (controller *ReportTaskController) Delete(c *gin.Context) {
	zap.S().Info("ReportTaskController->Delete")
	name := c.Param("name")
	if controller.service.Delete(name) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除上报服务失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除上报服务成功！", "data": nil})
}

func (controller *ReportTaskController) Find(c *gin.Context) {
	zap.S().Info("ReportTaskController->Find")
	name := c.Param("name")
	reportTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找上报服务失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找上报服务成功！", "data": reportTask})
}

// FindAll 获取所有上报服务数据
func (controller *ReportTaskController) FindAll(c *gin.Context) {
	zap.S().Info("ReportTaskController->FindAll")
	reportTasks, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有上报服务数据失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有上报服务数据成功！", "data": reportTasks})
}

func (controller *ReportTaskController) Start(c *gin.Context) {
	zap.S().Info("ReportTaskController->Start")
	name := c.Param("name")
	reportTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "找不到相关服务:" + name, "data": nil})
		return
	}

	reportTask.Status = 1

	// service.update内置开启/停止任务
	err = controller.service.Update(&reportTask)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改任务状态失败:" + name, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "启动任务成功！", "data": reportTask})
}

func (controller *ReportTaskController) Stop(c *gin.Context) {
	zap.S().Info("ReportTaskController->Stop")
	name := c.Param("name")
	reportTask, err := controller.service.Find(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "找不到相关服务:" + name, "data": nil})
		return
	}

	reportTask.Status = 0

	// service.update内置开启/停止任务
	err = controller.service.Update(&reportTask)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改任务状态失败:" + name, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "停止任务成功！", "data": reportTask})
}
