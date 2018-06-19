package routers

import (
	"github.com/Liv1020/move-car/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/register", controllers.User.Register)
	}
}
