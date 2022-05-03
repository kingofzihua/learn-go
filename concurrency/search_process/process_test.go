package main

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestProcess(t *testing.T) {
	err := process("hello world")
	fmt.Println(err)
}

func TestProcessTimeout(t *testing.T) {
	defer goleak.VerifyNone(t)
	result, err := processTimeout("hello world", 1000*time.Millisecond)
	fmt.Println(result)
	fmt.Println(err)
	time.Sleep(2000 * time.Millisecond)
}
