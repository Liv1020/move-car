package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	mpoauth "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

type user struct{}

// User 用户
var User = user{}

// Create 注册
func (t *user) Create(c *gin.Context) {
	form := new(form)
	err := c.BindJSON(form)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	db := components.App.DB()

	qr := new(models.Qrcode)
	if err := db.Where("id = ?", form.QrCodeID).Last(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	u := new(models.User)
	if err := db.Where("id = ?", "openid").Last(u).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			components.ResponseError(c, 1, err)
			return
		}
	}

	u.OpenID = "openid"
	u.Mobile = form.Mobile
	u.PlateNumber = form.PlateNumber
	if err := db.Save(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	qr.UserID = u.ID
	if err := db.Save(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, u)
}

// OAuth OAuth
func (t user) OAuth(c *gin.Context) {
	conf := components.App.Config

	callUrl := conf.Wechat.OAuthUrl
	authUrl := mpoauth.AuthCodeURL(conf.Wechat.AppID, callUrl, "snsapi_base", "STATE")

	c.Redirect(301, authUrl)
}

// Code Code
func (t user) Code(c *gin.Context) {
	code := c.GetString("code")
	conf := components.App.Config

	p := mpoauth.NewEndpoint(conf.Wechat.AppID, conf.Wechat.AppSecret)
	cli := &oauth2.Client{
		Endpoint: p,
	}

	token, err := cli.ExchangeToken(code)
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

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&info)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	c.SetCookie("wechat", string(b), int(token.ExpiresIn), "", "", false, false)
	c.Redirect(301, "http://localhost:8080")
}

type form struct {
	QrCodeID    int    `json:"qr_code_id"`
	Mobile      string `json:"mobile"`
	PlateNumber string `json:"plate_number"`
}
