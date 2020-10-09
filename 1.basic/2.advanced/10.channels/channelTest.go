package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Println("写入：", i)
		ch <- i
	}
}

func readFromChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		v := <-ch
		fmt.Println("读取：", v)
	}
}

func main() {
	var ch1 chan int            // 声明一个通道
	ch1 = make(chan int)        // 未初始化的通道不能存储数据，初始化一个通道
	ch2 := make(chan string, 2) // 声明并初始化一个带缓冲空间的通道

	// 通过匿名函数向通道中写入数据，通过 <- 方式写入
	go func() { ch1 <- 1 }()
	go func() { ch2 <- `a` }()

	v1 := <-ch1 // 从通道中读取数据
	v2 := <-ch2
	fmt.Println(v1) // 输出：1
	fmt.Println(v2) // 输出：a

	// 写入，读取通道数据
	ch3 := make(chan int, 1) // 初始化一个带缓冲空间的通道
	go readFromChannel(ch3)
	go writeToChannel(ch3)

	// 主线程休眠1秒，让出执行权限给子 Go 程，即通过 go 开启的 goroutine，不然主程序会直接结束
	time.Sleep(1 * time.Second)
}
