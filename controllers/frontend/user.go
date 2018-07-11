package frontend

import (
	"errors"

	"time"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/Liv1020/move-car-api/models"
	"github.com/Liv1020/move-car-api/resources"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	mpOauth "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	mpUser "gopkg.in/chanxuehong/wechat.v2/mp/user"
)

type user struct{}

// User 用户
var User = user{}

// Update Update
func (t *user) Update(c *gin.Context) {
	auth := middlewares.JwtAuthFromClaims(c)

	form := new(form)
	err := c.BindJSON(form)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if err := form.validate(); err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	db := components.App.DB()

	qr := new(models.Qrcode)
	if err := db.Where("id = ?", form.QrCode).Last(qr).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	u := new(models.User)
	if err := db.Where("id = ?", auth.ID).Last(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

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

// IsSubscribe IsSubscribe
func (t *user) IsSubscribe(c *gin.Context) {
	auth := middlewares.JwtAuthFromClaims(c)
	db := components.App.DB()

	u := new(models.User)
	if err := db.Where("openid = ?", auth.OpenID).Last(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	wu, err := mpUser.Get(components.App.WechatClient(), auth.OpenID, mpOauth.LanguageZhCN)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	u.IsSubscribe = wu.IsSubscriber
	if err := db.Save(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, resources.NewUser(u))
}

type form struct {
	QrCode      string `json:"qr_code"`
	Mobile      string `json:"mobile"`
	PlateNumber string `json:"plate_number"`
	Code        string `json:"code"`
}

func (t *form) validate() error {
	if t.Mobile == "" {
		return errors.New("手机号码不能为空")
	}
	if t.Code == "" {
		return errors.New("验证码不能为空")
	}
	db := components.App.DB()
	sc := new(models.SmsCode)
	if err := db.Where("mobile = ? AND expired_at > ? AND is_valid = 0", t.Mobile, time.Now()).Last(sc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("请重新发送验证")
		}

		return err
	}
	if sc.Code != t.Code {
		return errors.New("验证码错误")
	}
	sc.IsValid = 1
	if err := db.Save(sc).Error; err != nil {
		return err
	}
	if t.PlateNumber == "" {
		return errors.New("车牌号不能为空")
	}
	if t.QrCode == "" {
		return errors.New("二维码不能为空")
	}

	return nil
}

// Confirm Confirm
func (t *user) Confirm(c *gin.Context) {
	auth := middlewares.JwtAuthFromClaims(c)
	db := components.App.DB()

	type form struct {
		Wait int `json:"wait"`
	}

	p := new(form)
	if err := c.BindJSON(p); err != nil {
		return
	}

	now := time.Now()
	mt := now.Add(time.Duration(p.Wait) * time.Minute)
	auth.MoveAt = &mt
	if err := db.Save(auth).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	WS.SendWait(auth)

	components.ResponseSuccess(c, nil)
}
