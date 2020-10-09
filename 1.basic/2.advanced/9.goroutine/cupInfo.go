package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func calc() {
	for {
		i++
	}
}

func main() {

	/*
	 cpu 的核心数
	*/
	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)

	//设置 cpu，能够把 内核都用上，跑满cpu 性能
	// 设置4，只跑一半
	runtime.GOMAXPROCS(cpu)

	for i := 0; i < 10; i++ {
		go calc()
	}

	time.Sleep(time.Hour)
}
