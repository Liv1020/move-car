package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterMiddleware RegisterMiddleware
func RegisterMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8010", "http://mc.liv1020.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}
