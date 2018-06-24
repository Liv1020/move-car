package models

import "github.com/jinzhu/gorm"

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
}
