package models

// User User
type User struct {
	ID     int `gorm:"primary_key"`
	Mobile int `gorm:"column:mobile"`
}
