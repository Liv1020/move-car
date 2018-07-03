package main

import (
	"fmt"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/Liv1020/move-car-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	middlewares.RegisterMiddleware(r)

	routers.RegisterRouter(r)

	r.Run(fmt.Sprintf(":%d", components.App.Config().Port))
}
