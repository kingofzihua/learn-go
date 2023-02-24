package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
