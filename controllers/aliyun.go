package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/models"
	"github.com/denverdino/aliyungo/sms"
	"github.com/gin-gonic/gin"
)

type aliyun struct {
}

// Aliyun Aliyun
var Aliyun = aliyun{}

// Call Call
func (t *aliyun) Call(c *gin.Context) {
	components.ResponseSuccess(c, nil)
}

// Sms Sms
func (t *aliyun) Sms(c *gin.Context) {
	conf := components.App.Config.Aliyun
	db := components.App.DB()

	cli := sms.NewDYSmsClient(conf.AccessKeyId, conf.AccessKeySecret)
	type form struct {
		Mobile string `json:"mobile"`
	}

	f := new(form)
	c.BindJSON(f)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%04v", rnd.Int31n(10000))

	res, err := cli.SendSms(&sms.SendSmsArgs{
		PhoneNumbers:  f.Mobile,
		SignName:      "阿里云短信测试专用",
		TemplateCode:  "SMS_137875084",
		TemplateParam: `{"code": "` + code + `"}`,
	})
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if res.Code != "OK" {
		components.ResponseError(c, 1, errors.New(res.Message))
		return
	}

	now := time.Now()
	sc := &models.SmsCode{
		Code:      code,
		Mobile:    f.Mobile,
		IsValid:   0,
		ExpiredAt: now.Add(time.Minute * 30),
	}
	if err := db.Save(sc).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	components.ResponseSuccess(c, nil)
}
