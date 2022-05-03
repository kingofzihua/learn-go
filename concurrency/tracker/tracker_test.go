package main

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestNewTracker(t *testing.T) {
	defer goleak.VerifyNone(t)
	tracker := NewTracker(10)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("tracker run panic:", err)
			}
		}()
		tracker.Run()
	}()

	_ = tracker.Event(context.Background(), "event 1")
	_ = tracker.Event(context.Background(), "event 2")
	_ = tracker.Event(context.Background(), "event 3")

	// tracker 退出的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := tracker.ShutDown(ctx)

	fmt.Println("ShutDown err:", err)

}
