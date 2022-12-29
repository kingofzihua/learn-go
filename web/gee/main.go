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

	eng.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	log.Fatal(eng.Run(":8080"))
}
