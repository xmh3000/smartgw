package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartgw/api/domain"
	"smartgw/api/service"
	"smartgw/lib/web/middleware"
)

type AccountController struct {
	userService service.UserService
}

func NewAccountController(userService service.UserService) *AccountController {
	return &AccountController{
		userService: userService,
	}
}

func (a *AccountController) Login(c *gin.Context) {
	user := domain.User{}
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "1",
			"message": "用户数据解析失败！",
			"data":    "",
		})
		return
	}

	if a.userService.Valid(user.Username, user.Password) {
		if token, err := middleware.GenerateToken(user.Username); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0",
				"message": "登录成功！",
				"data":    token,
			})
		}
	}
}
