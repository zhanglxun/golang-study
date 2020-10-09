package main

import "fmt"

//默认标准写法，入参和返回值
func add(x int, y int) int {
	return x + y
}

func main() {

	fmt.Println(add(42, 13))

}
