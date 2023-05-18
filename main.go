package main

import (
	"errors"
	"fmt"
	"go-artisan/cache"
	"go-artisan/route"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	//router()
	cacheRun()
}

func router() {
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

func cacheRun() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, errors.New(key + " not exist")
		},
	))

	addr := "localhost:9999"
	peers := cache.NewHTTPPool(addr)
	log.Println("gee cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
