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

	qr := &models.Qrcode{
		UserID: 0,
	}
	if err := db.Save(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, qr)
}
