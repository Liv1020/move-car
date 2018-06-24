package main

import (
	"fmt"

	"github.com/Liv1020/move-car/components"
	"github.com/Liv1020/move-car/middlewares"
	"github.com/Liv1020/move-car/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	middlewares.RegisterMiddleware(r)
	routers.RegisterRouter(r)

	r.Run(fmt.Sprintf(":%d", components.App.Config.Port))
}
