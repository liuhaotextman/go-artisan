package route

import (
	"fmt"
	"net/http"
)

type HandlerFunction func(*Context)

type router struct {
	routers map[string]HandlerFunction
}

func New() *router {
	return &router{
		routers: make(map[string]HandlerFunction),
	}
}

func (r *router) GET(uri string, handler HandlerFunction) {
	r.addRoute("GET", uri, handler)
}

func (r *router) POST(uri string, handler HandlerFunction) {
	r.addRoute("POST", uri, handler)
}

func (r *router) PUT(uri string, handler HandlerFunction) {
	r.addRoute("PUT", uri, handler)
}

func (r *router) DELETE(uri string, handler HandlerFunction) {
	r.addRoute("DELETE", uri, handler)
}

func (r *router) addRoute(method string, uri string, handler HandlerFunction) {
	key := method + ":" + uri
	r.routers[key] = handler
}

func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + ":" + request.URL.Path
	handler, ok := r.routers[key]
	if !ok {
		_, err := fmt.Fprintf(writer, "uri not found")
		if err != nil {
			return
		}
		return
	}
	context := NewContext(writer, request)
	handler(context)
}

func (r *router) Run(address string) error {
	err := http.ListenAndServe(address, r)
	return err
}
