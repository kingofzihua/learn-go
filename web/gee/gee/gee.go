package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	// 前缀
	prefix string
	// 支持嵌套
	parent *RouterGroup
	// 所有分组共享 Engine 实例
	engine *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // 所有的分组
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	// New RouterGroup
	rg := &RouterGroup{engine: engine}
	engine.RouterGroup = rg
	engine.groups = []*RouterGroup{rg}

	return engine
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	eg := rg.engine

	newGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: eg,
	}

	eg.groups = append(eg.groups, newGroup)

	return newGroup
}

func (rg *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := rg.prefix + comp

	// log info
	log.Printf("Route %4s - %s register", method, pattern)

	rg.engine.router.addRoute(method, pattern, handler)
}

func (rg *RouterGroup) Get(path string, handler HandlerFunc) {
	rg.addRoute("GET", path, handler)
}

func (rg *RouterGroup) Post(path string, handler HandlerFunc) {
	rg.addRoute("POST", path, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}
