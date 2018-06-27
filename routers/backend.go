package routers

import (
	bc "github.com/Liv1020/move-car/controllers/backend"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterBackend 注册路由
func registerBackend(router *gin.Engine) {
	router.POST("/backend/user/login", bc.Admin.Login)

	backend := router.Group("/backend")
	backend.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{

	}
}
