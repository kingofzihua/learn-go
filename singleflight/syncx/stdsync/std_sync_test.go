package stdsync

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"testing"
	"time"
)

var g singleflight.Group

func TestExclusiveCallDo(T *testing.T) {
	ticker := time.NewTicker(time.Microsecond * 1000)
	var num int32

	for ; ; <-ticker.C {

		go func() {
			atomic.AddInt32(&num, 1)
			_, err := getDataFromDb("")
			if err != nil {
				panic(err)
			}
			//fmt.Println(num, res)
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
