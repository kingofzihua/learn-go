package main

import (
	"log"
	"net/http"

	"github.com/kingofzihua/learn-go/web/gee/gee"
)

func main() {
	eng := gee.New()

	eng.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q \n", c.Path)
	})

	eng.Get("/header", func(c *gee.Context) {
		c.Json(http.StatusOK, c.Req.Header)
	})

	eng.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s , you're at %s \n", c.Query("name"), c.Path)
	})

	eng.Get("/hi/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Param("name"))
	})

	eng.Get("/hi/a/b", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q \n", c.Path)
	})

	eng.Get("/hi/a/c", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q \n", c.Path)
	})

	eng.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	eng.Get("/assets/*filepath", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	v1 := eng.Group("/v1")
	{
		v1.Use(gee.Logger())
		user := v1.Group("/user")
		{
			user.Get("/:name", func(c *gee.Context) {
				c.Json(http.StatusOK, gee.H{"name": c.Param("name")})
			})
		}
		order := v1.Group("/order/:order_sn")
		{
			order.Get("/detail", func(c *gee.Context) {
				c.Json(http.StatusOK, gee.H{"order_sn": c.Param("order_sn")})
			})
		}
	}

	log.Fatal(eng.Run(":8080"))
}
