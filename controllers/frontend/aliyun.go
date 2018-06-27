package frontend

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/components/aliyun/vms"
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
	type params struct {
		QrCode string `json:"qr_code"`
	}

	p := new(params)
	c.BindJSON(p)

	db := components.App.DB()

	row := new(models.Qrcode)
	if err := db.Preload("User").Where("id = ?", p.QrCode).Last(row).Error; err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	// wx := components.App.Config().Wechat
	// co := components.App.WechatClient()
	// message := &template.TemplateMessage{
	// 	ToUser:     row.User.OpenID,
	// 	TemplateId: wx.TemplateID,
	// 	URL:        wx.ConfirmUrl,
	// 	Data:       []byte(`{}`),
	// }
	// _, err := template.Send(co, message)
	// if err != nil {
	// 	components.ResponseError(c, 1, err)
	// 	return
	// }

	conf := components.App.Config().Aliyun
	cli := vms.NewDYVmsClient(conf.AccessKeyId, conf.AccessKeySecret)
	res, err := cli.SendVms(&vms.SendVmsArgs{
		CalledShowNumber: "09314267618",
		CalledNumber:     row.User.Mobile,
		TtsCode:          "TTS_137689626",
		TtsParam:         `{"plate":"` + row.User.PlateNumber + `"}`,
		Volume:           100,
		PlayTimes:        3,
	})
	if err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if res.Code != "OK" {
		components.ResponseError(c, 1, errors.New(res.Message))
		return
	}

	components.ResponseSuccess(c, nil)
}

// Sms Sms
func (t *aliyun) Sms(c *gin.Context) {
	conf := components.App.Config().Aliyun
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
		SignName:      conf.Vms.SignName,
		TemplateCode:  conf.Vms.TemplateCode,
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
