package main

import (
	"fmt"
	"sync"
)

// 线程安全的计数器类型
type CounterMore struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *CounterMore) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *CounterMore) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	var counterMore CounterMore
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counterMore.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counterMore.Count())
}
