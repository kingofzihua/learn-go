package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	e.router.addRoute(method, path, handler)
}

func (e *Engine) Get(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

func (e *Engine) Post(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}
