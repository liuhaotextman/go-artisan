package main

import (
	"fmt"
	"go-artisan/route"
)

func main() {
	router := route.New()
	router.GET("/", func(c *route.Context) {
		c.HTML(200, "this is sucks")
	})
	router.GET("index", func(context *route.Context) {
		context.HTML(200, "home page")
	})
	router.GET("detail/:id", func(context *route.Context) {
		context.HTML(200, "detail id="+context.Params("id"))
	})

	adminGroup := router.Group("admin")
	adminGroup.GET("login", func(context *route.Context) {
		context.HTML(200, "admin user login page")
	})
	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err.Error())
	}
}
