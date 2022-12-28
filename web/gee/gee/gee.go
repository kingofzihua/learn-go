package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	e.router[key] = handler
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

	key := req.Method + "-" + req.URL.Path

	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
