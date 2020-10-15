package main

import (
	"fmt"
	"sync"
)

/**
	Mutex 会嵌入到其它 struct 中使用，比如下面的方式
    采用嵌入字段的方式。
	通过嵌入字段，你可以在这个 struct 上直接调用 Lock/Unlock 方法。
*/
type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {

	var counter Counter
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count)

}
