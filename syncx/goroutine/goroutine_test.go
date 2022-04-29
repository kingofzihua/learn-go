package goroutine

import "testing"

func TestSubGoroutinePanic(t *testing.T) {
	Go(func() {
		panic("sub goroutine panic!")
	})
}
