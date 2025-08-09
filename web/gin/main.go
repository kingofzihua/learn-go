package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// /h el lo /gin
// /h el lo 2/gin
// /h el o :username /name
// /h i:usernamess

func main() {
	router := gin.Default()

	router.GET("/index", Index)
	router.GET("/hello/gin", HelloGin)
	router.GET("/hello2/gin", HelloGin)
	// middleware
	// 全局中间件
	router.Use(CorsMiddleware())
	/**

	/user/gordon              match
	/user/you                 match
	/user/gordon/profile      no match
	/user/                    no match

	*/
	router.GET("/user/:username", User)
	router.GET("/hi:username", User)        // 允许
	router.GET("/helo:username/Name", User) // 允许
	//router.GET("/:d-:m-:y/name", User)    // 不允许

	//router.GET("/blog/:/", BlogIndex) // params 路由必须有名字
	// router.GET("/blog/index", BlogIndex) // 同一个路由前缀，不允许同时定义 static 和 params

	/**
	/blog/go/request-routers            match: category="go", post="request-routers"
	/blog/go/request-routers/           no match, but the router would redirect
	/blog/go/                           no match
	/blog/go/request-routers/comments   no match
	*/
	router.GET("/blog/:category/:post/", Blog)
	router.GET("/blog/:category/:post/:name", Blog)
	router.GET("/bg/:category/:post/:name", Blog)

	router.GET("/testing/:name", Testing)

	router.GET("/m3u", M3U)

	/**
	/static/                     match
	/static/index.html           match
	/static/js/test.js           match
	*/
	router.StaticFS("/static/", http.Dir("./public"))

	v1 := router.Group("/api/v1")
	v1.Use(AuthMiddleware()) // v1 分组的中间件
	{
		// /api/v1/users
		v1.GET("/users", GetUser) // 继承 CorsMiddleware, AuthMiddleware

		// /api/v1/users/create
		// 这个路由有自己的特定中间件 LoggerMiddleware
		v1.POST("/users/create", LoggerMiddleware(), CreateUser) // 继承 CorsMiddleware, AuthMiddleware, 并使用 LoggerMiddleware
	}

	log.Fatal(router.Run(":8080"))
}
