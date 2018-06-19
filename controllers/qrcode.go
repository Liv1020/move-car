package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
)

type qrcode struct{}

// QrCode 二维码
var QrCode = qrcode{}

// Create 创建
func (t *qrcode) Create(c *gin.Context) {
	db := components.App.DB()

	row := &models.Qrcode{
		UserID: 0,
	}
	if err := db.Save(row).Error; err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
		"data":    row,
	})
}
