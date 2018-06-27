package routers

import (
	"github.com/Liv1020/move-car/controllers/frontend"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/gin-gonic/gin"
)

// registerFrontend 注册路由
func registerFrontend(router *gin.Engine) {
	f := router.Group("/frontend")
	{
		f.GET("/ws", frontend.WS.Handle)

		wechat := f.Group("/wechat")
		{
			wechat.POST("/oauth", frontend.Wechat.Oauth)
			wechat.POST("/server", frontend.Wechat.Server)
		}

		user := f.Group("/user")
		user.Use(middlewares.JwtMiddleware.MiddlewareFunc())
		{
			user.POST("/update", frontend.User.Update)
		}

		qr := f.Group("/qrcode")
		qr.Use(middlewares.JwtMiddleware.MiddlewareFunc())
		{
			qr.GET("/view", frontend.QrCode.View)
			qr.POST("/create", frontend.QrCode.Create)
			qr.GET("/search", frontend.QrCode.Search)
		}

		js := f.Group("/js")
		js.Use(middlewares.JwtMiddleware.MiddlewareFunc())
		{
			js.GET("/config", frontend.JS.Config)
		}

		aliyun := f.Group("/aliyun")
		aliyun.Use(middlewares.JwtMiddleware.MiddlewareFunc())
		{
			aliyun.POST("/call", frontend.Aliyun.Call)
			aliyun.POST("/sms", frontend.Aliyun.Sms)
		}
	}
}
