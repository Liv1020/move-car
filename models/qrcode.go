package models

import "github.com/jinzhu/gorm"

// Qrcode Qrcode
type Qrcode struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`

	User *User
}
