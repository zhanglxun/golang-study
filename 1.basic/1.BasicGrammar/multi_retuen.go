package main

import "fmt"

/*
多返回值,可以支持一个函数内，两个活多个返回值，定义好返回值的顺序，return 既可。
*/

func add(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {

	sum, sub := add(7, 3)
	fmt.Println(sum, sub)
}
