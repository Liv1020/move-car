package backend

import (
	"errors"

	"fmt"
	"time"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/gin-gonic/gin"
)

type admin struct {
}

// Admin Admin
var Admin = admin{}

// Login Login
func (*admin) Login(c *gin.Context) {
	type form struct {
		Username string
		Password string
	}

	f := new(form)
	if err := c.BindJSON(f); err != nil {
		components.ResponseError(c, 1, err)
		return
	}

	if f.Username != "admin" {
		components.ResponseError(c, 1, errors.New("账号未找到"))
		return
	}

	if f.Password != "admin123456" {
		components.ResponseError(c, 1, errors.New("密码错误"))
		return
	}

	type token struct {
		Token     string `json:"token"`
		ExpiredAt int    `json:"expired_at"`
	}

	now := time.Now()
	uid := 1

	tk := new(token)
	tk.Token = middlewares.JwtMiddleware.TokenGenerator(fmt.Sprintf("%d", uid))
	tk.ExpiredAt = int(now.Add(2*time.Hour - 30).Unix())

	components.ResponseSuccess(c, tk)
}
