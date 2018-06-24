package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// SmsCode SmsCode
type SmsCode struct {
	gorm.Model
	Mobile    string
	Code      string
	IsValid   int
	ExpiredAt time.Time

	User *User
}
