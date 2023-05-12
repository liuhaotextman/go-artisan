package route

import "strings"

type GroupRouter struct {
	prefix     string
	parent     *GroupRouter
	middleware []HandlerFunction
	router     *router
}

func (g *GroupRouter) GET(uri string, handler HandlerFunction) {
	g.addRoute("GET", uri, handler)
}

func (g *GroupRouter) POST(uri string, handler HandlerFunction) {
	g.addRoute("POST", uri, handler)
}

func (g *GroupRouter) PUT(uri string, handler HandlerFunction) {
	g.addRoute("PUT", uri, handler)
}

func (g *GroupRouter) DELETE(uri string, handler HandlerFunction) {
	g.addRoute("DELETE", uri, handler)
}

func (g *GroupRouter) addRoute(method string, uri string, handler HandlerFunction) {
	pattern := g.prefix + uri
	g.router.addRoute(method, pattern, handler)
}

func (g *GroupRouter) Group(uri string) *GroupRouter {
	if strings.HasPrefix(uri, "/") {
		uri = uri[0 : len(uri)-1]
	}
	if !strings.HasSuffix(uri, "/") {
		uri = uri + "/"
	}
	return &GroupRouter{prefix: uri, parent: g, router: g.router}
}
