package frontend

import (
	"fmt"

	"time"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	mpoauth "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

type oauth struct{}

// Oauth 用户
var Oauth = oauth{}

// Code Code
func (t oauth) Code(c *gin.Context) {
	type params struct {
		Code string
	}

	form := new(params)
	c.BindJSON(form)

	conf := components.App.Config()

	p := mpoauth.NewEndpoint(conf.Wechat.AppID, conf.Wechat.AppSecret)
	cli := &oauth2.Client{
		Endpoint: p,
	}

	token, err := cli.ExchangeToken(form.Code)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	info, err := mpoauth.GetUserInfo(token.AccessToken, token.OpenId, mpoauth.LanguageZhCN, nil)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	db := components.App.DB()

	u := new(models.User)
	if err := db.Where("openid = ?", info.OpenId).Last(u).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			components.ResponseError(c, 1, err)
			return
		}
	}

	u.OpenID = info.OpenId
	u.Nickname = info.Nickname
	u.Sex = info.Sex
	u.City = info.City
	u.Province = info.Province
	u.Country = info.Country
	u.HeadImageUrl = info.HeadImageURL
	if err := db.Save(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	now := time.Now()
	at := new(appToken)
	at.Token = middlewares.JwtMiddleware.TokenGenerator(fmt.Sprintf("%d", u.ID))
	at.ExpiredAt = int(now.Add(2*time.Hour - 30).Unix())
	at.User = u

	components.ResponseSuccess(c, at)
}

type appToken struct {
	Token     string       `json:"token"`
	ExpiredAt int          `json:"expired_at"`
	User      *models.User `json:"user"`
}
