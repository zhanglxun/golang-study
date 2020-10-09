package main

import "fmt"

func main() {

	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 1)
	fmt.Printf("c=%v\n", c)
	//把100 入队，到队列里面
	c <- 100

	fmt.Printf("chan c=%v\n", c)

	data := <-c
	fmt.Printf("data:%v\n", data)

	//<-c
}
