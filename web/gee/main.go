package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kingofzihua/learn-go/web/gee/gee"
)

func main() {
	eng := gee.New()

	eng.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
	})

	eng.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header [%q] = %q \n", k, v)
		}
	})

	log.Fatal(eng.Run(":8080"))
}
