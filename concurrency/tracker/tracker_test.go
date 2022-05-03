package main

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"sync"
	"testing"
	"time"
)

func TestNewTracker(t *testing.T) {
	defer goleak.VerifyNone(t)

	tracker := NewTracker(1)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("tracker run panic:", err)
			}
		}()
		tracker.Run()
	}()

	wg.Wait()

	var err error
	err = tracker.Event(context.Background(), "event 1")
	if err != nil {
		fmt.Println(err)
	}

	err = tracker.Event(context.Background(), "event 2")
	if err != nil {
		fmt.Println(err)
	}

	err = tracker.Event(context.Background(), "event 3")
	if err != nil {
		fmt.Println(err)
	}

	// tracker 退出的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = tracker.ShutDown(ctx)
	if err != nil {
		fmt.Println("ShutDown err:", err)
	}

}
