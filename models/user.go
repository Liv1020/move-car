package models

import "github.com/jinzhu/gorm"

const (
	// SUBSCRIBE_NO 未关注
	SUBSCRIBE_NO = 0
	// SUBSCRIBE_YES 已关注
	SUBSCRIBE_YES = 1
)

// User User
type User struct {
	gorm.Model
	OpenID       string `gorm:"column:openid"`
	Nickname     string `gorm:"column:nickname"`
	Sex          int    `gorm:"column:sex"`
	City         string `gorm:"column:city"`
	Province     string `gorm:"column:province"`
	Country      string `gorm:"column:country"`
	HeadImageUrl string `gorm:"column:head_image_url"`
	Mobile       string `gorm:"column:mobile"`
	PlateNumber  string `gorm:"column:plate_number"`
	IsSubscribe  int    `gorm:"column:is_subscribe"`
	WaitMinute   int    `gorm:"column:wait_minute"`
}
