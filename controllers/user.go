package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/gin-gonic/gin"
)

type user struct {
}

// User 用户
var User = user{}

// Register 注册
func (t *user) Register(c *gin.Context) {
	components.DB.Where("")

	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
