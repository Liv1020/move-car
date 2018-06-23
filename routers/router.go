package routers

import (
	"github.com/Liv1020/move-car/controllers"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	auth := router.Group("oauth")
	{
		auth.GET("/index", controllers.Oauth.Index)
		auth.GET("/code", controllers.Oauth.Code)
	}

	router.GET("/")

	user := router.Group("/user")
	user.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		user.POST("/create", controllers.User.Create)
	}

	qr := router.Group("/qrcode")
	qr.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		qr.POST("/create", controllers.QrCode.Create)
	}

	js := router.Group("/js")
	js.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		js.GET("/config", controllers.JS.Config)
	}

	aliyun := router.Group("/aliyun")
	aliyun.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		aliyun.POST("/call", controllers.Aliyun.Call)
	}
}
