package zeromicro

import "sync"

// copy from https://github.com/zeromicro/go-zero/blob/master/core/syncx/singleflight.go
type (
	// SingleFlight lets the concurrent calls with the same key to share the call result.
	// For example, A called F, before it's done, B called F. Then B would not execute F,
	// and shared the result returned by F which called by A.
	// The calls with the same key are dependent, concurrent calls share the returned values.
	// A ------->calls F with key<------------------->returns val
	// B --------------------->calls F with key------>returns val

	// SingleFlight 定义接口，有2个方法 Do 和 DoEx，其实逻辑是一样的，DoEx 多了一个标识，主要看Do的逻辑就够了
	SingleFlight interface {
		Do(key string, fn func() (interface{}, error)) (interface{}, error)
		DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
	}

	// 定义 call 的结构
	call struct {
		wg  sync.WaitGroup // 用于实现通过1个 call，其他 call 阻塞
		val interface{}    // 表示 call 操作的返回结果
		err error          // call 操作发生的错误
	}

	// 总控结构，实现 SingleFlight 接口
	flightGroup struct {
		calls map[string]*call
		lock  sync.Mutex
	}
)

// NewSingleFlight returns a SingleFlight.
func NewSingleFlight() SingleFlight {
	return &flightGroup{
		calls: make(map[string]*call),
	}
}

func (g *flightGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	// 对 key 发起 call 请求， 如果此时已经有其他协程已经在发起 call 请求就阻塞住（done 为 true 的情况），等待拿到结果后直接返回。
	c, done := g.createCall(key)
	if done { // 如果 done 是 false，说明当前协程是第一个发起 call 的协程
		return c.val, c.err
	}

	// 真正地发起 call 请求（
	g.makeCall(c, key, fn)
	return c.val, c.err
}

func (g *flightGroup) DoEx(key string, fn func() (interface{}, error)) (val interface{}, fresh bool, err error) {
	c, done := g.createCall(key)
	if done {
		return c.val, false, c.err
	}

	g.makeCall(c, key, fn)
	return c.val, true, c.err
}

func (g *flightGroup) createCall(key string) (c *call, done bool) {
	g.lock.Lock()
	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c, true
	}

	c = new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	return c, false
}

// 发起请求
func (g *flightGroup) makeCall(c *call, key string, fn func() (interface{}, error)) {
	defer func() {
		g.lock.Lock()
		delete(g.calls, key)
		g.lock.Unlock()
		c.wg.Done()
	}()

	c.val, c.err = fn()
}
