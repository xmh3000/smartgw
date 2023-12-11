package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"smartgw/api/service"
	"smartgw/lib/web/middleware"
)

// AccountController 账号相关操作
type AccountController struct {
	service service.UserService
}

func NewAccountController(service service.UserService) *AccountController {
	return &AccountController{
		service: service,
	}
}

// Login 系统登录
func (controller *AccountController) Login(c *gin.Context) {
	zap.S().Info("AccountController->Login")

	request := middleware.LoginRequest{}
	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "1",
			"message": "用户数据解析失败",
			"data":    "",
		})
		// 错误返回
		return
	}
	zap.S().Info(request)
	if controller.service.Valid(request.Username, request.Password) {
		// 成功返回
		token, err := middleware.GenerateToken(request.Username)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    "1",
				"message": err.Error(),
				"data":    "",
			})
		}
		zap.S().Debugf("jwtToken[%v]", token)
		// 封装一个响应数据,返回用户名和token
		data := middleware.LoginResponse{
			Token:    token,
			Username: request.Username,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    "0",
			"message": "login sucess",
			"data":    data,
		})
		return
	}
}
