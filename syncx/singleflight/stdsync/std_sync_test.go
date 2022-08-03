package singleflight

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var g Group

func TestExclusiveCallDo(T *testing.T) {
	var num int32

	for num < 10 {
		go func() {
			atomic.AddInt32(&num, 1)
			_, err := getDataFromDb("")
			if err != nil {
				panic(err)
			}
		}()
	}
}

func getDataFromDb(name string) (res string, err error) {
	result, err, _ := g.Do(name, func() (interface{}, error) {
		fmt.Println("get buy db")
		time.Sleep(time.Millisecond * 1000)
		return "kingofzihua", nil
	})
	res = result.(string)
	return
}
