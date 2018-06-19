package components

import "github.com/jinzhu/gorm"

// DB DB
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer DB.Close()
}
