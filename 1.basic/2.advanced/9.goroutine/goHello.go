package main

import (
	"fmt"
	"time"
)

func PrintHello(i int) {
	fmt.Println("hello goroutine", i)

}

func main() {

	for i := 0; i < 10; i++ {
		go PrintHello(i)
	}
	time.Sleep(time.Second)
}
