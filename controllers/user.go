package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
)

type user struct{}

// User 用户
var User = user{}

// Create 注册
func (t *user) Create(c *gin.Context) {
	auth := components.GetAuthFromClaims(c)

	form := new(form)
	err := c.BindJSON(form)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	db := components.App.DB()

	qr := new(models.Qrcode)
	if err := db.Where("id = ?", form.QrCodeID).Last(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	u := new(models.User)
	if err := db.Where("id = ?", auth.ID).Last(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	u.Mobile = form.Mobile
	u.PlateNumber = form.PlateNumber
	if err := db.Save(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	qr.UserID = u.ID
	if err := db.Save(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, u)
}

type form struct {
	QrCodeID    int    `json:"qr_code_id"`
	Mobile      string `json:"mobile"`
	PlateNumber string `json:"plate_number"`
}
