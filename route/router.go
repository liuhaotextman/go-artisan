package route

import (
	"net/http"
	"strings"
)

type HandlerFunction func(*Context)

type router struct {
	roots   map[string]*node
	routers map[string]HandlerFunction
	groups  []*GroupRouter
	*GroupRouter
}

func New() *router {
	router := &router{
		routers: make(map[string]HandlerFunction),
		roots:   make(map[string]*node),
	}
	router.GroupRouter = &GroupRouter{
		router: router,
	}
	router.groups = []*GroupRouter{router.GroupRouter}
	return router
}

func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	vs := strings.Split(pattern, "/")
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, uri string, handler HandlerFunction) {
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(uri)
	r.roots[method].insert(uri, parts, 0)
	key := method + "-" + uri
	r.routers[key] = handler
}

func (r *router) getRoute(method, pattern string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := parsePattern(pattern)
	node := root.search(searchParts, 0)
	params := make(map[string]string)
	if node != nil {
		parts := parsePattern(node.pattern)
		for key, part := range parts {
			if strings.HasPrefix(part, ":") {
				params[part[1:]] = searchParts[key]
			}

			if strings.HasPrefix(part, "*") {
				params[part[1:]] = strings.Join(searchParts[key:], "/")
				break
			}
		}
	}

	return node, params
}

func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var middlewares []HandlerFunction
	for _, group := range r.groups {
		if strings.HasPrefix(request.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	context := NewContext(writer, request)
	context.handlers = middlewares
	node, params := r.getRoute(context.method, context.uri)
	if node == nil {
		context.JSON(400, map[string]string{
			"code": "400",
			"msg":  "url not found",
		})
		return
	}
	key := context.method + "-" + node.pattern
	handler, ok := r.routers[key]
	if !ok {
		context.JSON(400, map[string]string{
			"code": "400",
			"msg":  "url handler not found",
		})
		return
	}
	context.params = params
	context.handlers = append(context.handlers, handler)
	context.Next()
}

func (r *router) Run(address string) error {
	err := http.ListenAndServe(address, r)
	return err
}
