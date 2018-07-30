package routers

import (
	"github.com/Liv1020/move-car-api/controllers/backend"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterBackend 注册路由
func registerBackend(router *gin.Engine) {
	router.POST("/backend/user/login", backend.Admin.Login)

	b := router.Group("/backend")
	b.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		user := b.Group("/user")
		{
			user.GET("/search", backend.User.Search)
		}
	}
}
