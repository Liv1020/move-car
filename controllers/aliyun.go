package controllers

import (
	"github.com/Liv1020/move-car/components"
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
