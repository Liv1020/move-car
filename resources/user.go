package resources

import "github.com/Liv1020/move-car/models"

// User User
type User struct {
	ID           uint   `json:"id"`
	OpenID       string `json:"open_id"`
	Nickname     string `json:"nickname"`
	Sex          int    `json:"sex"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	HeadImageUrl string `json:"head_image_url"`
	Mobile       string `json:"mobile"`
	PlateNumber  string `json:"plate_number"`
	IsSubscribe  int    `json:"is_subscribe"`
}

// NewUser NewUser
func NewUser(u *models.User) *User {
	user := &User{
		ID:           u.ID,
		OpenID:       u.OpenID,
		Nickname:     u.Nickname,
		Sex:          u.Sex,
		City:         u.City,
		Province:     u.Province,
		Country:      u.Country,
		HeadImageUrl: u.HeadImageUrl,
		Mobile:       u.Mobile,
		PlateNumber:  u.PlateNumber,
		IsSubscribe:  u.IsSubscribe,
	}

	return user
}
