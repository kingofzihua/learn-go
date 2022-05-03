package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// searchSleep 是一个耗时的操作， 我需要给他设置 timeout 如果处理时间超过了 100 毫秒就返回 timeout error

func main() {
	result, err := processTimeout("hello world", 100*time.Millisecond)
	fmt.Println(result, err)
	time.Sleep(200 * time.Millisecond)
}

func searchSleep(term string) (string, error) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("search sleep return:" + term)
	return "some value:" + term, nil
}

type result struct {
	record string
	err    error
}

func processTimeout(term string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan result)

	// 超时以后的这个 goroutine 去哪里了？
	go func() {
		defer func() {
			fmt.Println("goroutine exit") // defer 也未执行
		}()
		record, err := searchSleep(term)

		fmt.Println("search return")

		// 使用 select ， 当超时当时候
		select {
		case <-ctx.Done():
			fmt.Println("goroutine ctx done")
			return
		case ch <- result{record, err}:
			fmt.Println("ch get data") // 超时以后并未打印
			return
		}
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("search time out")
	case result := <-ch:
		if result.err != nil {
			return "", result.err
		}
		return result.record, nil
	}
}
