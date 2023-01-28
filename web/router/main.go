package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/index", Index)
	router.GET("/hello/gin", HelloGin)
	router.GET("/hi/gin", HelloGin)
	/**

	/user/gordon              match
	/user/you                 match
	/user/gordon/profile      no match
	/user/                    no match

	*/
	router.GET("/user/:username", User)

	//router.GET("/blog/:/", BlogIndex) // params 路由必须有名字
	// router.GET("/blog/index", BlogIndex) // 同一个路由前缀，不允许同时定义 static 和 params

	/**
	/blog/go/request-routers            match: category="go", post="request-routers"
	/blog/go/request-routers/           no match, but the router would redirect
	/blog/go/                           no match
	/blog/go/request-routers/comments   no match
	*/
	router.GET("/blog/:category/:post", Blog)
	router.GET("/blog/:category/:post/:name", Blog)
	router.GET("/b/:category/:post/:name", Blog)

	/**
	/static/                     match
	/static/index.html           match
	/static/js/test.js           match
	*/
	router.ServeFiles("/static/*filepath", http.Dir("./public"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func HelloGin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello gin")
}

func BlogIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "blog index")
}

func Blog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "category:%s, post:%s!\n", ps.ByName("category"), ps.ByName("post"))
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func User(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("username"))
}
