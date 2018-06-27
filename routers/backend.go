package routers

import (
	bc "github.com/Liv1020/move-car/controllers/backend"
	"github.com/Liv1020/move-car/controllers/frontend"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterBackend 注册路由
func registerBackend(router *gin.Engine) {
	router.POST("/backend/user/login", bc.Admin.Login)

	backend := router.Group("/backend")
	backend.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		qr := backend.Group("/qrcode")
		{
			qr.GET("/view", frontend.QrCode.View)
			qr.POST("/create", frontend.QrCode.Create)
			qr.GET("/search", frontend.QrCode.Search)
		}
	}
}
