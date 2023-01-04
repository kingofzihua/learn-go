package gee

import (
	"reflect"
	"testing"
)

func newTestRoute() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func Test_parsePattern(t *testing.T) {
	tests := []struct {
		pattern string
		want    []string
	}{
		{
			pattern: "/p/:name",
			want:    []string{"p", ":name"},
		},
		{
			pattern: "/p/*",
			want:    []string{"p", "*"},
		},
		{
			pattern: "/p/*name/*",
			want:    []string{"p", "*name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := parsePattern(tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router_getRoute(t *testing.T) {

	tests := []struct {
		method  string
		path    string
		pattern string
		params  map[string]string
	}{
		{
			method:  "GET",
			path:    "/hello/geektutu",
			pattern: "/hello/:name",
			params:  map[string]string{"name": "geektutu"},
		},
		{
			method:  "GET",
			path:    "/hello/b/c",
			pattern: "/hello/b/c",
			params:  make(map[string]string),
		},
		{
			method:  "GET",
			path:    "/assets/file1.txt",
			pattern: "/assets/*filepath",
			params:  map[string]string{"filepath": "file1.txt"},
		},
		{
			method:  "GET",
			path:    "/assets/css/test.css",
			pattern: "/assets/*filepath",
			params:  map[string]string{"filepath": "css/test.css"},
		},
	}

	r := newTestRoute()

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			node, params := r.getRoute(tt.method, tt.path)
			if node == nil {
				t.Errorf("getRoute() node is nil mathod = %v, path %v", tt.method, tt.path)
			}
			if !reflect.DeepEqual(node.pattern, tt.pattern) {
				t.Errorf("getRoute() pattern = %v, want %v", node.pattern, tt.pattern)
			}

			if (params != nil || tt.params != nil) && !reflect.DeepEqual(params, tt.params) {
				t.Errorf("getRoute() got = %v, want %v", params, tt.params)
			}
		})
	}
}
