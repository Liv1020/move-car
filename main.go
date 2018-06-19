package main

import (
	"github.com/Liv1020/move-car/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	routers.RegisterRouter(r)

	r.Run(":8080")
}
