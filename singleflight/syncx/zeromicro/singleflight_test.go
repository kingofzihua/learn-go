package zeromicro

import (
	"fmt"
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
