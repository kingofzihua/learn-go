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

	/**
	/static/                     match
	/static/index.html           match
	/static/js/test.js           match
	*/
	router.StaticFS("/static/", http.Dir("./public"))

	log.Fatal(router.Run(":8080"))
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}

func HelloGin(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}

func BlogIndex(c *gin.Context) {
	c.String(http.StatusOK, "blog index")
}

func Blog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"category": c.Param("category"), "post": c.Param("post")})
}

type Person struct {
	Name    string `uri:"name"`
	Address string `form:"address"`
}

func Testing(c *gin.Context) {
	var person Person

	err := c.ShouldBind(&person)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "address": person.Address})
}

func User(c *gin.Context) {
	c.String(http.StatusOK, "hello, %s!\n", c.Param("username"))
}
