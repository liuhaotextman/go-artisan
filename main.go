package main

import (
	"fmt"
	"go-artisan/route"
	"net/http"
)

func main() {
	router := route.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "this is worked")
	})
	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err.Error())
	}
}
