package main

import "fmt"

// defer 定义不会立即执行，会在函数 return时候执行
// defer 会把定义 压到 堆栈 里面，执行的时候，反着返回，先进后出
// 用来在程序最后 执行 回收资源的操作，是 比较好的一个使用场景

// defer 定义不会立即执行，会在函数 return时候执行
// defer 会把定义 压到 堆栈 里面，执行的时候，反着返回，先进后出
// 用来在程序最后 执行 回收资源的操作，是 比较好的一个使用场景
func testDefer() {
	defer fmt.Println("hello")
	defer fmt.Println("jarvis")
	defer fmt.Println("zhang")

	fmt.Println("123")
	fmt.Println("456")
}
func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// 求 1，100 的所有质数的 ，打印到屏幕上
func examlple1() {
	for i := 1; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("%d is prime \n", i)
		}
	}
}

func main() {

	testDefer()

	examlple1()

}
