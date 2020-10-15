package main

import (
	"fmt"
	"sync"
)

/**
  go run -race geekTimeMuTex.go

  通过过race 工具，能方便的检测 并发中，读写操作的问题

*/
func main() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)

}
