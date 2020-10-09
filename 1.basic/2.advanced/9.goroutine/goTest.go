package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {

	go say("list")
	say("1")

	fmt.Println("i get exited.")
	time.Sleep(5 * time.Second)
}
