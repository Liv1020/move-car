package middlewares

import (
	"time"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// JwtMiddleware JwtMiddleware
var JwtMiddleware *jwt.GinJWTMiddleware
var JwtHttpMiddleware *jwt.GinJWTMiddleware

// JwtAuthFromClaims JwtAuthFromClaims
func JwtAuthFromClaims(c *gin.Context) *models.User {
	user := new(models.User)
	val, ok := c.Get("auth")
	if !ok {
		return user
	}

	if user, ok = val.(*models.User); !ok {
		return user
	}

	return user
}

func init() {
	// JwtMiddleware
	JwtMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "Move Car",
		Key:        []byte("b24bd75e0ab"),
		Timeout:    2 * time.Hour,
		MaxRefresh: 2 * time.Hour,
		Authorizator: func(userID string, c *gin.Context) bool {
			db := components.App.DB()
			user := new(models.User)
			if err := db.Where("id = ?", userID).Model(&models.User{}).Last(user).Error; err != nil {
				return false
			}

			c.Set("auth", user)

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status":  code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
	}

	// JwtHttpMiddleware
	JwtHttpMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "Move Car",
		Key:        []byte("b24bd75e0ab"),
		Timeout:    2 * time.Hour,
		MaxRefresh: 2 * time.Hour,
		Authorizator: func(userID string, c *gin.Context) bool {
			db := components.App.DB()
			user := new(models.User)
			if err := db.Where("id = ?", userID).Model(&models.User{}).Last(user).Error; err != nil {
				return false
			}

			c.Set("auth", user)

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status":  code,
				"message": message,
			})
		},
		TokenLookup: "query:token",
	}
}
