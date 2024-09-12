package main

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"log"
	_ "net"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)
	defer cancelFunc()
	go func(ctx context.Context) {
		for {
			fmt.Println("go for")
			select {
			case <-ctx.Done():
				log.Println("done")
				return
			default:
				for {
				}
			}
		}
	}(ctx)

	<-ctx.Done()

	fmt.Println("end")
}
