package main

import (
	"database/sql"
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
	//cacheRun()
	myOrm()
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

func myOrm() {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() {
		_ = db.Close()
	}()

	_, _ = db.Exec("drop table if exists User;")
	_, _ = db.Exec("create table User(Name text);")
	result, err := db.Exec("insert into User(`name`) values (?), (?)")
	if err != nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("select Name from User limit 1;")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
