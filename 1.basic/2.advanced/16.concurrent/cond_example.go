package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知，阻塞当前goroutine
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠1秒，使所有goroutine进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1秒后下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1秒后下发下一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1) // 睡眠1秒，等待所有goroutine执行完毕
}

//条件变量和锁结合使用，在并发时如果逻辑不严谨容易发生死锁，
//所以尽量不要使用条件变量，推荐用 sync.WaitGroup 来实现并发时 Go 程间的同步。
