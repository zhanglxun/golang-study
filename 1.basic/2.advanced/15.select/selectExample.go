package main

import (
	"fmt"
	"time"
)

//select是用来同步两个线程之间的数据

/*
A.同时监听一个或多个channel，直到其中一个channel ready
B. 如果其中多个channel同时ready，随机选择一个进行操作。
C. 语法和switch case有点类似，代码可读性更好。
*/

func server1(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "response from server1"
}

func server2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "response from server2"
}

func main() {

	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	/*
		s1 := <-output1
		fmt.Println("s1:", s1)

		s2 := <-output2
		fmt.Println("s2:", s2)

	*/
	select {
	case s1 := <-output1:
		fmt.Println("s1", s1)
	case s2 := <-output2:
		fmt.Println("s2", s2)
	default:
		fmt.Println("run default")
	}

}
