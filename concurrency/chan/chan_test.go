package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// chan close 所有阻塞的消费者都能收到消息 val, ok := <- chan  ok = true: 正常 ; ok = false : chan close
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("goroutine 1 exit")
			wg.Done()
		}()
		val, ok := <-stop
		fmt.Println("val=", val, " ok=", ok)
	}()
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("goroutine 2 exit")
			wg.Done()
		}()
		<-stop
	}()

	go func() {
		time.Sleep(2 * time.Second)
		stop <- struct{}{}
		time.Sleep(2 * time.Second)
		close(stop)
	}()

	wg.Wait()
}
