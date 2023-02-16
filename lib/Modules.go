package lib

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"smartgw/lib/config"
	"smartgw/lib/database"
	"smartgw/lib/logger"
	"smartgw/lib/web"
)

var Modules = fx.Options(
	config.Module,
	logger.Module,
	database.Module,
	web.Module,

	// 初始化web服务器
	fx.Invoke(func(gin *gin.Engine) {}),
)
