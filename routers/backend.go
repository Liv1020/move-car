package routers

import (
	"github.com/gin-gonic/gin"
)

// RegisterBackend 注册路由
func registerBackend(router *gin.Engine) {
	backend := router.Group("/backend")
	{
		user := backend.Group("/user")
		{
			user.POST("/login")
		}
	}
}
