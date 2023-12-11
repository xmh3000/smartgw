package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
)

// UserController 用户相关控制器，对本系统而言，没有什么作用，只是一个示例
type UserController struct {
	service service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}

// Add 新增用户
func (controller *UserController) Add(c *gin.Context) {
	zap.S().Info("UserController->Add")
	user := domain.User{}
	// 解析用户数据
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析用户数据失败！", "data": nil})
		return
	}

	// 新增用户失败
	if controller.service.Add(&user) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "新增用户失败！", "data": nil})
		return
	}

	// 新增用户成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "新增用户成功！", "data": user})
}

func (controller *UserController) Update(c *gin.Context) {
	zap.S().Info("UserController->Update")
	user := domain.User{}
	// 解析用户数据
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "解析用户数据失败！", "data": nil})
		return
	}

	// 修改用户失败
	if controller.service.Update(&user) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "修改用户失败！", "data": nil})
		return
	}

	// 修改用户成功
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改用户成功！", "data": user})
}

func (controller *UserController) Delete(c *gin.Context) {
	zap.S().Info("UserController->Delete")

	username := c.Param("name")

	if controller.service.Delete(username) != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "删除用户失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除用户成功！", "data": nil})
}

func (controller *UserController) Find(c *gin.Context) {
	zap.S().Info("UserController->Find")

	username := c.Param("name")

	user, err := controller.service.Find(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "查找用户失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "查找用户成功！", "data": user})
}

// FindAll 获取所有用户数据
func (controller *UserController) FindAll(c *gin.Context) {
	zap.S().Info("UserController->FindAll")

	users, err := controller.service.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "获取所有用户数据失败！", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "获取所有用户数据成功！", "data": users})
}
