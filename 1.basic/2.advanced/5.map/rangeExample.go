package main

import "fmt"

func testRange() {

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

// 双赋值监测 某个键是否存在
func testDoubleGet() {

}

func main() {
	// 获取range 的值返回
	testRange()
}
