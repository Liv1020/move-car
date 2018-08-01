package backend

import (
	"strconv"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/models"
	"github.com/gin-gonic/gin"
)

type user struct{}

// User 用户
var User = user{}

// Search Search
func (t *user) Search(c *gin.Context) {
	db := components.App.DB().Model(&models.User{})
	list := &list{
		Data:  make([]*u, 0, 15),
		Total: 0,
	}
	p, _ := strconv.Atoi(c.Query("page"))
	s, _ := strconv.Atoi(c.Query("size"))
	page := components.Page{
		Page: p,
		Size: s,
	}

	if err := db.Offset(page.GetOffset()).Limit(page.GetLimit()).Count(&list.Total).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if list.Total == 0 {
		components.ResponseSuccess(c, list)
		return
	}

	var rows []*models.User
	if err := db.Offset(page.GetOffset()).Limit(page.GetLimit()).Find(&rows).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	for _, row := range rows {
		list.Data = append(list.Data, &u{
			ID:          row.ID,
			Nickname:    row.Nickname,
			City:        row.City,
			Mobile:      row.Mobile,
			PlateNumber: row.PlateNumber,
		})
	}

	components.ResponseSuccess(c, list)
}

type list struct {
	Data  []*u `json:"data"`
	Total int  `json:"total"`
}

type u struct {
	ID          uint   `json:"id"`
	Nickname    string `json:"nickname"`
	City        string `json:"city"`
	Mobile      string `json:"mobile"`
	PlateNumber string `json:"plate_number"`
}
