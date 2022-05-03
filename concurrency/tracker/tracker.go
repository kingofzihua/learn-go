package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

//  cap 表示最大可以存储多少个 event
func NewTracker(cap int) *Tracker {
	return &Tracker{
		ch:   make(chan string),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println("event default")
		return errors.New(fmt.Sprintf("event chan full:%s", data))
	}
}

func (t *Tracker) Run() {
	fmt.Println("tracker run")

	for data := range t.ch {
		time.Sleep(2 * time.Second)
		fmt.Println(data)
	}

	t.stop <- struct{}{}
}

func (t *Tracker) ShutDown(ctx context.Context) error {
	close(t.ch)
	select {
	case <-t.stop:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
