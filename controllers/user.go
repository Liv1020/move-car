package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type user struct{}

// User 用户
var User = user{}

// Create 注册
func (t *user) Create(c *gin.Context) {
	form := new(form)
	err := c.BindJSON(form)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err,
		})
		return
	}

	db := components.App.DB()

	qr := new(models.Qrcode)
	if err := db.Where("id = ?", form.QrCodeID).Last(qr).Error; err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err,
		})
		return
	}

	u := new(models.User)
	if err := db.Where("id = ?", "openid").Last(u).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(200, gin.H{
				"status":  1,
				"message": err,
			})
			return
		}
	}

	u.OpenID = "openid"
	u.Mobile = form.Mobile
	u.PlateNumber = form.PlateNumber
	if err := db.Save(u).Error; err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err,
		})
		return
	}

	qr.UserID = u.ID
	if err := db.Save(qr).Error; err != nil {
		c.JSON(200, gin.H{
			"status":  1,
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
		"data":    u,
	})
}

type form struct {
	QrCodeID    int    `json:"qr_code_id"`
	Mobile      string `json:"mobile"`
	PlateNumber string `json:"plate_number"`
}
