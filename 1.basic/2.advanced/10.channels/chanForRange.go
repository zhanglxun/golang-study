package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {

	ch := make(chan int, 3)

	go producer(ch)

	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("receive:", v)
		time.Sleep(2 * time.Second)
	}

}
