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
	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err.Error())
	}
}
