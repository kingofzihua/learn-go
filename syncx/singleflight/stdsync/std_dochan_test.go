package singleflight

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func getData(ctx context.Context, key string) (string, error) {
	ch := g.DoChan(key, func() (ret interface{}, err error) {
		fmt.Println("do chan")
		time.Sleep(3 * time.Second)
		return "kingofzihua", nil
	})

	select {
	case <-ctx.Done():
		return "", errors.New("timeout")
	case res := <-ch:
		return res.Val.(string), nil
	}

}

func TestDoChan(t *testing.T) {
	defer goleak.VerifyNone(t)
	var count = 10
	var wg sync.WaitGroup

	ct, _ := context.WithTimeout(context.TODO(), 2*time.Second)

	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			res, err := getData(ct, "kkk")
			fmt.Println(res, err)
		}()
	}

	wg.Wait()
	time.Sleep(5 * time.Second)
}
