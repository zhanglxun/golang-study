package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("run ", i, " times")
	}
	fmt.Println("calc finish.")
}

//通过使用 go calc() 来调用协程
func main() {
	go calc()
	fmt.Println("i exited.")
	time.Sleep(11 * time.Second)
}
