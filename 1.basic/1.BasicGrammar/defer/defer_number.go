package main

import "fmt"

func test() int {
	//临时的变量
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func test2() (j int) {

	j = 0
	defer func() {
		j += 1
		fmt.Println("defer2")
	}()
	return j
}

func main() {
	fmt.Println("return", test())

	fmt.Println("return", test2())
}

//返回值为：
//defer2
//defer1
//return 0
//执行 return 语句后，Go 会创建一个临时变量保存返回值，
//因此，defer 语句修改了局部变量 i，并没有修改返回值。

/**
对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存.
因此，defer 语句修改了 i，即对返回值产生了影响。
*/
