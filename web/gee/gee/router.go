package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 只允许一个 * 并且在最后
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)

			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// insert node
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)

	// insert handler
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	sp := parsePattern(pattern)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(sp, 0)
	if n == nil {
		return nil, nil
	}

	parts := parsePattern(n.pattern) // 选中的节点

	for i, part := range parts {
		// 遇到 :
		// route => /:name  path =>/value
		// name = value
		if part[0] == ':' {
			params[part[1:]] = sp[i]
		}

		// 遇到 * ，将后面所有的都当作 一个字符串处理
		// route => /*name  path =>/value/123123
		// name = value/123123
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(sp[i:], "/")
			break
		}
	}

	return n, params
}

func (r *router) handle(c *Context) {

	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		// handler
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
