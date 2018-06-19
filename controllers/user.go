package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
)

type user struct {
}

// User 用户
var User = user{}

// Register 注册
func (t *user) Register(c *gin.Context) {
	db := components.App.DB

	row := new(models.User)
	if err := db.Where("id = ?", 1).Last(row).Error; err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
		"row":     row,
	})
}
