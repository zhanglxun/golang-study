package main

import (
	"fmt"
	"strings"
	"time"
)

/**
闭包的示例
闭包是由函数及其相关引用环境组合而成的实体
*/

// 闭包，定义一个局部变量，函数作为返回值
func Adder() func(int) int {
	var x int
	//匿名函数，作为返回值
	return func(d int) int {
		x += d
		return x
	}
}
func testClosure1() {
	f := Adder()
	ret := f(1)
	fmt.Printf("ret=%d\n", ret)
	ret = f(20)
	fmt.Printf("ret=%d\n", ret)
	ret = f(100)
	fmt.Printf("ret=%d\n", ret)

	//重新调用，由从0 开始
	f1 := Adder()
	rem := f1(1)
	fmt.Printf("ret=%d\n", rem)

}

// 闭包，把局部变量定义的，前面例子，作为参数，传入进来
func addClosure(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClosure2() {
	tmp1 := addClosure(10)
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := addClosure(100)
	fmt.Println(tmp2(1), tmp2(2))
}

// 闭包，字符串的相关操作 例子
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testClosure3() {
	f3 := makeSuffix(".bmp")
	f4 := makeSuffix(".jpg")

	fmt.Println(f3("test.bmp"))
	fmt.Println(f4("test"))
}

// 闭包，多返回值的一个 闭包例子
func multiNum(base int) (func(int) int, func(int) int) {
	addMulti := func(i int) int {
		base += i
		return base
	}
	subMulti := func(i int) int {
		base -= i
		return base
	}

	return addMulti, subMulti
}
func testmultiNum() {
	f1, f2 := multiNum(10)

	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))

}

func main() {
	//闭包的基本使用
	testClosure1()

	testClosure2()

	testClosure3()

	testmultiNum()

	//闭包的副作用
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("-----")
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}
