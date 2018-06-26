package controllers

import (
	"time"

	"fmt"

	"bytes"
	"math/rand"
	"strings"

	"github.com/Liv1020/move-car/components"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
)

type js struct{}

// JS 二维码
var JS = js{}

// Config Config
func (t *js) Config(c *gin.Context) {
	wechat := components.App.Config().Wechat
	now := time.Now()
	conf := &config{
		AppId:     wechat.AppID,
		Timestamp: fmt.Sprintf("%d", now.Unix()),
		NonceStr:  randomString(12, "0aA"),
		Signature: "",
	}

	as := core.NewDefaultAccessTokenServer(wechat.AppID, wechat.AppSecret, nil)

	clt := core.NewClient(as, nil)

	ts := jssdk.NewDefaultTicketServer(clt)

	ticket, err := ts.Ticket()
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	conf.Signature = jssdk.WXConfigSign(ticket, conf.NonceStr, conf.Timestamp, "")

	components.ResponseSuccess(c, conf)
}

func randomString(randLength int, randType string) (result string) {
	var num = "0123456789"
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}

type config struct {
	AppId     string `json:"app_id"`
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}
