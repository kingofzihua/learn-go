package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 多个 goroutine 等待 一个 wait group
func TestMultipleGoroutineWaitSingleGroup(T *testing.T) {

	var wg sync.WaitGroup

	var out sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		// 等待两秒
		time.Sleep(time.Second * 2)
		fmt.Println("wg done")
	}()

	for i := 0; i < 10; i++ {
		out.Add(1)
		go func(i int) {
			defer out.Done()
			fmt.Println("wait result start", i)
			wg.Wait()
			fmt.Println("wait result end", i)
		}(i)
	}

	out.Wait()
}
