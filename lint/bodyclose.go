package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://example.com/") // Wrong case
	if err != nil {
		// handle error
	}
	body, err := ioutil.ReadAll(resp.Body)
}
