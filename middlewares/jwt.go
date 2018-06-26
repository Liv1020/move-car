package middlewares

import (
	"time"

	"errors"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// JwtMiddleware JwtMiddleware
var JwtMiddleware *jwt.GinJWTMiddleware

// JwtAuthFromClaims JwtAuthFromClaims
func JwtAuthFromClaims(c *gin.Context) *models.User {
	claims := jwt.ExtractClaims(c)
	uid := claims["id"]

	db := components.App.DB()

	user := new(models.User)
	db.Where("id = ?", uid).Last(user)

	return user
}

func init() {
	// the jwt middleware
	JwtMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "Move Car",
		Key:        []byte("b24bd75e0ab"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authorizator: func(userID string, c *gin.Context) bool {
			db := components.App.DB()
			count := 0
			if err := db.Where("id = ?", userID).Model(&models.User{}).Count(&count).Error; err != nil {
				return false
			}

			if count == 0 {
				return false
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			components.ResponseError(c, code, errors.New(message))
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
