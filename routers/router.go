package routers

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) {
	registerFrontend(router)
	registerBackend(router)
}
