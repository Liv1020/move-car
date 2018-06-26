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
		auth.POST("/code", controllers.Oauth.Code)
	}

	user := router.Group("/user")
	user.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		user.POST("/update", controllers.User.Update)
	}

	qr := router.Group("/qrcode")
	qr.Use(middlewares.JwtMiddleware.MiddlewareFunc())
	{
		qr.GET("/view", controllers.QrCode.View)
		qr.POST("/create", controllers.QrCode.Create)
		qr.GET("/search", controllers.QrCode.Search)
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
		aliyun.POST("/sms", controllers.Aliyun.Sms)
	}
}
