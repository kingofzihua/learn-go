package errors

import (
	"fmt"
	"testing"
	"time"
)

// 无意义 因为 recover 执行的时候并没有 panic ，所以没有错误，
// panic 的时候 就挂掉了
func TestRecoverInMain(t *testing.T) {
	if err := recover(); err != nil {
		fmt.Printf("recover panic:%v\n", err)
	}

	panic("goroutine panic ")
}

// 可以正确捕获 panic 后会执行 defer ， recover 的时候因为之前 panic 了所以能拿到 recover
func TestRecoverInMainDefer(t *testing.T) {
	defer func() {
		fmt.Println("first defer")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover panic:%v\n", err)
		}
	}()

	panic("goroutine panic ")
}

// 主 goroutine 并不能捕获其他的 goroutine 的panic
// 如果子goroutine 没有自己捕获异常，则 就会导致整个进程挂掉
func TestMainGoroutineDeferSubGoroutinePanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover panic:%v\n", err)
		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		panic("goroutine panic ")
	}()

	time.Sleep(time.Second * 100)
}

// 可以捕获 panic
// 子 goroutine 自己捕获 panic 可以保证进程正常运行
func TestSubGoroutineDeferPanic(t *testing.T) {
	go func() {
		ticker := time.NewTicker(time.Microsecond * 500)

		for range ticker.C {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Printf("recover panic:%v\n", err)
					}
				}()
				time.Sleep(time.Second * 1)

				panic("goroutine panic ")
			}()

		}
	}()

	time.Sleep(time.Second * 2)

}
