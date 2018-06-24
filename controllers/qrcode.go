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

// View View
func (t *qrcode) View(c *gin.Context) {
	db := components.App.DB()
	qrID := c.Query("qr")

	qr := new(models.Qrcode)
	if err := db.Where("id = ?", qrID).Last(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	out := &qrCode{
		ID:     qr.ID,
		UserID: qr.UserID,
	}

	components.ResponseSuccess(c, out)
}

type qrCode struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
}
