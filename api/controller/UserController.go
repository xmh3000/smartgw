package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/service"
)

type UserController struct {
	service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService,
	}
}

func (u *UserController) Find(c *gin.Context) {
	zap.S().Info("UserController->Find")
	username := c.Param("name")

	if user, err := u.UserService.Find(username); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "查找用户失败！",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "查找用户成功！",
			"data":    user,
		})
	}
}
