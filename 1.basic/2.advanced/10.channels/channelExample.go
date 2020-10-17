package main

import "fmt"

func main() {

	var c chan int
	// 还没有初始化，为nil 值
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 1)
	//初始化了，值分配了地址
	fmt.Printf("c=%v\n", c)

	//把100 入队，到队列里面
	c <- 100

	//这个时候，赋值了，但是 chan 也只是一个通道，还是一个地址
	fmt.Printf("chan c=%v\n", c)

	data := <-c
	// 把值输出，才会得到 100
	fmt.Printf("data:%v\n", data)

	//<-c
}
