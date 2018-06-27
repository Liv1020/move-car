package backend

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type qrcode struct{}

// QrCode 二维码
var QrCode = qrcode{}

// Search Search
func (t *qrcode) Search(c *gin.Context) {
	type search struct {
		components.Page
	}

	type list struct {
		Data  []*models.Qrcode `json:"data"`
		Total int              `json:"total"`
	}

	s := new(search)
	c.ShouldBindWith(s, binding.JSON)

	db := components.App.DB()

	out := new(list)
	if err := db.Model(&models.Qrcode{}).Count(&out.Total).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if out.Total == 0 {
		components.ResponseSuccess(c, out)
		return
	}

	if err := db.Preload("User").Offset(s.GetOffset()).Limit(s.GetLimit()).Find(&out.Data).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, out)
}

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

	out := &qrCode{
		ID:     qr.ID,
		UserID: qr.UserID,
	}
	components.ResponseSuccess(c, out)
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
