package main

import (
	"fmt"
	"math"
)

//判断奇数偶数
func testfi() {
	num := 9

	if num%2 == 0 {
		fmt.Printf("number is enen \n")
	} else {
		fmt.Printf("number is odd \n")
	}

	//定义的变量，可以写一行
	if number := 100; number%2 == 0 {
		fmt.Printf("100 is enen \n")
	} else {
		fmt.Printf("100 is odd \n")
	}
}

// 注意 fallthrough 的使用，可以穿透到下一个case 的执行，只穿透下一个。
func testSwitch() {

	switch a := 1; a {
	case 1:
		fmt.Printf("a=1 \n")
		fallthrough
	case 2:
		fmt.Printf("a=2 \n")
	case 3:
		fmt.Printf("a=3 \n")
	default:
		fmt.Printf("nothing to do")
	}
}

//乘法口诀表打印
func testMulti() {
	for i := 0; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d =%d \t", j, j, j*i)
		}
		fmt.Println()
	}
}

//数据处理
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	testfi()
	testSwitch()

	testMulti()

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
