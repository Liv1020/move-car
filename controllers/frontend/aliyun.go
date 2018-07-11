package frontend

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/components/aliyun/vms"
	"github.com/Liv1020/move-car-api/models"
	"github.com/denverdino/aliyungo/sms"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
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

	wx := components.App.Config().Wechat
	co := components.App.WechatClient()

	now := time.Now()
	message := &template.TemplateMessage{
		ToUser:     row.User.OpenID,
		TemplateId: wx.TemplateID,
		URL:        wx.ConfirmUrl,
		Data:       []byte(`{"first":{"value":"兰智挪车让爱车更智能","color":"#173177"},"keyword1":{"value":"` + row.User.PlateNumber + `","color":"#173177"},"keyword2":{"value":"` + now.Format("2006-01-02 15:04:05") + `","color":"#173177"},"remark":{"value":"因你目前停放车子的位置影响了其他人出行，请尽快前往挪车！【点击前往】","color":"#173177"}}`),
	}
	_, err := template.Send(co, message)
	if err != nil {
		components.App.Logger().Errorf("发送模板消息错误：%s", err)
	}

	time.AfterFunc(time.Second*30, func() {
		log := components.App.Logger()
		aliyun := components.App.Config().Aliyun
		cli := vms.NewDYVmsClient(aliyun.AccessKeyId, aliyun.AccessKeySecret)
		res, err := cli.SendVms(&vms.SendVmsArgs{
			CalledShowNumber: aliyun.Vms.CalledShowNumber,
			CalledNumber:     row.User.Mobile,
			TtsCode:          aliyun.Vms.TtsCode,
			TtsParam:         `{"plate":"` + row.User.PlateNumber + `"}`,
			Volume:           100,
			PlayTimes:        3,
		})
		if err != nil {
			log.Errorf("发送模板消息错误：%s", err)
			return
		}

		if res.Code != "OK" {
			log.Errorf("发送模板消息错误：%s", res.Message)
			return
		}
	})

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
		SignName:      conf.Sms.SignName,
		TemplateCode:  conf.Sms.TemplateCode,
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
