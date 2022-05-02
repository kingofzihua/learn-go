package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(stop)
	}()

	go func() {
		handler := http.NewServeMux()

		handler.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello debug!")
		})

		done <- serve(":8081", handler, stop)
	}()

	go func() {
		handle := http.DefaultServeMux

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello World")
		})
		done <- serve(":8080", handle, stop)
	}()

	fmt.Printf("error : %+v", <-done)
	fmt.Printf("error : %+v", <-done)

	//var stopping bool // chan 只能 close 1 次

	//for i := 0; i < cap(done); i++ {
	//	fmt.Println("<- done ")
	//
	//	if err := <-done; err != nil {
	//		fmt.Printf("error : %+v", err)
	//	}
	//	fmt.Println("stopping==", stopping)
	//	if !stopping {
	//		stopping = true
	//		close(stop)
	//	}
	//}

}

func serve(addr string, handler http.Handler, stop chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	time.Sleep(2 * time.Second)

	return fmt.Errorf("serve %s is error:%w\n", addr, s.ListenAndServe())
}
