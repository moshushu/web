package routes

import (
	"net/http"
	"web/logger"
	"web/settings"

	"github.com/gin-gonic/gin"
)

func Routes(mode string) *gin.Engine {
	// 出厂模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	// 不使用gin默认的中间件，使用定制logger中间件
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
