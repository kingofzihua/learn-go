package main

import (
	"context"
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
		ch:   make(chan string, cap),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
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
