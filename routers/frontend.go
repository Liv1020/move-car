package routers

import (
	"github.com/Liv1020/move-car-api/controllers/backend"
	"github.com/Liv1020/move-car-api/controllers/frontend"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/gin-gonic/gin"
)

// registerFrontend 注册路由
func registerFrontend(router *gin.Engine) {
	f := router.Group("/frontend")
	{
		f.GET("/ws", middlewares.JwtHttpMiddleware.MiddlewareFunc(), frontend.WS.Handle)

		wechat := f.Group("/wechat")
		{
			wechat.POST("/oauth", frontend.Wechat.Oauth)
			wechat.Any("/server", frontend.Wechat.Server)
		}

		user := f.Group("/user")
		user.Use(middlewares.JwtMiddleware.MiddlewareFunc())
		{
			user.POST("/is-subscribe", frontend.User.IsSubscribe)
			user.POST("/update", frontend.User.Update)
			user.POST("/confirm", frontend.User.Confirm)
			user.GET("/search", backend.User.Search)
		}

		qr := f.Group("/qrcode")
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

		k8s := f.Group("/k8s")
		{
			k8s.GET("/v1", frontend.K8S.V1)
			k8s.GET("/v2", frontend.K8S.V2)
		}
	}
}
