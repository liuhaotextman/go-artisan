package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	writer     http.ResponseWriter
	request    *http.Request
	method     string
	uri        string
	params     map[string]string
	statusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,
		method:  r.Method,
		uri:     r.URL.Path,
	}
}

func (c *Context) Header(key, value string) {
	c.writer.Header().Set(key, value)
}

func (c *Context) GetHeader(key string) string {
	return c.request.Header.Get(key)
}

func (c *Context) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c *Context) Params(key string) string {
	item, ok := c.params[key]
	if ok {
		return item
	}

	return ""
}

func (c *Context) PostForm(key string) string {
	return c.request.FormValue(key)
}

func (c *Context) Status(code int) {
	c.statusCode = code
	c.writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Header("content-type", "text/plain")
	c.Status(code)
	_, _ = c.writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(code int, html string) {
	c.Header("content-type", "text/html")
	c.Status(code)
	_, _ = c.writer.Write([]byte(html))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Header("content-type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}
