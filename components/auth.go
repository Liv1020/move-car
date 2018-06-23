package components

import (
	"github.com/Liv1020/move-car/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// GetAuthFromClaims GetAuthFromClaims
func GetAuthFromClaims(c *gin.Context) *models.User {
	claims := jwt.ExtractClaims(c)
	uid := claims["id"]

	db := App.DB()

	user := new(models.User)
	db.Where("id = ?", uid).Last(user)

	return user
}
