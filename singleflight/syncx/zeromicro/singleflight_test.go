package zeromicro

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestExclusiveCallDo(t *testing.T) {
	g := NewSingleFlight()
	v, err := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})
	if got, want := fmt.Sprintf("%v (%T)", v, v), "bar (string)"; got != want {
		t.Errorf("Do = %v; want %v", got, want)
	}
	if err != nil {
		t.Errorf("Do error = %v", err)
	}
}

func TestNewSingleFlight(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	round := 10
	var wg sync.WaitGroup
	barrier := NewSingleFlight()
	wg.Add(round)
	for i := 0; i < round; i++ {
		go func() {
			defer wg.Done()
			// 启用10个协程模拟获取缓存操作
			val, err := barrier.Do("get_rand_int", func() (interface{}, error) {
				time.Sleep(time.Second)
				return rand.Int(), nil
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}
	wg.Wait()
}

func TestMuCallDo(t *testing.T) {
	g := NewSingleFlight()
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			res, _ := g.Do("key", func() (interface{}, error) {
				fmt.Println(1231)
				time.Sleep(time.Second * 10)
				return "bar", nil
			})
			fmt.Println(res)
		}()
	}

	wg.Wait()

}
