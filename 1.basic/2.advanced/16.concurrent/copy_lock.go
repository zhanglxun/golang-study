package main

import (
	"fmt"
	"sync"
)

/** 是go vet 命令可以检查
go vet copy_lock.go
# command-line-arguments
./copy_lock.go:19:6: call of foo copies lock value: command-line-arguments.Counter1
./copy_lock.go:23:12: foo passes lock by value: command-line-arguments.Counter1
*/
type Counter1 struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter1
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter1) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
