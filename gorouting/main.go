package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var stop = make(chan int)
	go func() {
		fmt.Println("AAAAAAAA")
		for {
			select {
			case s := <-stop:
				fmt.Println("AA =>", s)
				return
			case <-time.After(3 * time.Second):
				fmt.Println("AA", time.Now().Format(time.DateTime))
			}
		}
	}()
	go func() {
		fmt.Println("BBBBBBB")
		for {
			select {
			case s := <-stop:
				fmt.Println("BB =>", s)
				return
			case <-time.After(3 * time.Second):
				fmt.Println("BB", time.Now().Format(time.DateTime))
			}
		}
	}()

	go func() {
		fmt.Println("CCCCC")
		for {
			select {
			case s := <-stop:
				fmt.Println("CC =>", s)
				return
			case <-time.After(3 * time.Second):
				fmt.Println("Cc", time.Now().Format(time.DateTime))
			}
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	close(stop)
	time.Sleep(1 * time.Second)
	fmt.Println("END=>")
}
