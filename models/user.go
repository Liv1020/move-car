package models

import "github.com/jinzhu/gorm"

// User User
type User struct {
	gorm.Model
	OpenID      string `gorm:"column:openid"`
	Mobile      string `gorm:"column:mobile"`
	PlateNumber string `gorm:"column:plate_number"`

	QrCodes []*Qrcode
}
