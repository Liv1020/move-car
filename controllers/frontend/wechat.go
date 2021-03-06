package frontend

import (
	"fmt"

	"time"

	"net/http"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/Liv1020/move-car-api/models"
	"github.com/Liv1020/move-car-api/resources"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	mpOauth "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	mpUser "gopkg.in/chanxuehong/wechat.v2/mp/user"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

type wechat struct{}

// Wechat Wechat
var Wechat = wechat{}

// Oauth Oauth
func (t *wechat) Oauth(c *gin.Context) {
	type params struct {
		Code string
	}

	form := new(params)
	c.BindJSON(form)

	conf := components.App.Config()

	p := mpOauth.NewEndpoint(conf.Wechat.AppID, conf.Wechat.AppSecret)
	cli := &oauth2.Client{
		Endpoint: p,
	}

	token, err := cli.ExchangeToken(form.Code)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	info, err := mpOauth.GetUserInfo(token.AccessToken, token.OpenId, mpOauth.LanguageZhCN, nil)
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	wu, err := mpUser.Get(components.App.WechatClient(), info.OpenId, mpOauth.LanguageZhCN)
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
	u.IsSubscribe = wu.IsSubscriber
	if err := db.Save(u).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	type appToken struct {
		Token     string          `json:"token"`
		ExpiredAt int             `json:"expired_at"`
		User      *resources.User `json:"user"`
	}

	now := time.Now()
	at := new(appToken)
	at.Token = middlewares.JwtMiddleware.TokenGenerator(fmt.Sprintf("%d", u.ID))
	at.ExpiredAt = int(now.Add(2*time.Hour - 600).Unix()) // 比微信的超时时间短10分钟
	at.User = resources.NewUser(u)

	components.ResponseSuccess(c, at)
}

// Server Server
func (*wechat) Server(c *gin.Context) {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(func(c *core.Context) {
		lg := components.App.Logger()
		lg.Infof("收到消息:\n%s\n", c.MsgPlaintext)

		c.NoneResponse()
	})
	mux.DefaultEventHandleFunc(func(c *core.Context) {
		lg := components.App.Logger()
		lg.Printf("收到事件:\n%s\n", c.MsgPlaintext)

		db := components.App.DB()

		switch c.MixedMsg.EventType {
		case "subscribe":
			u := new(models.User)
			if err := db.Where("openid = ?", c.MixedMsg.FromUserName).Last(u).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					c.RawResponse(err)
					return
				}
			}

			u.OpenID = c.MixedMsg.FromUserName
			u.IsSubscribe = models.SUBSCRIBE_YES
			if err := db.Save(u).Error; err != nil {
				c.RawResponse(err)
				return
			}
		}

		c.NoneResponse()
	})

	conf := components.App.Config().Wechat

	server := core.NewServer("", "", conf.Token, conf.EncodingAesKey, mux, &ErrorHandle{})
	server.ServeHTTP(c.Writer, c.Request, nil)
}

// ErrorHandle ErrorHandle
type ErrorHandle struct {
}

// ServeError ServeError
func (*ErrorHandle) ServeError(w http.ResponseWriter, r *http.Request, err error) {
	lg := components.App.Logger()
	lg.Infof("Wechat Server Error: %s", err)
}
