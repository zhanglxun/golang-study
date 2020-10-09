package main

import "fmt"

/*
 空接口 可以存储任意的类型
*/
func testInterface() {
	var a interface{}
	var b int = 100

	a = b
	fmt.Printf("%T, %v \n", a, a)

	var c string = "hello"
	a = c
	fmt.Printf("%T ,%v \n", a, a)

	var d map[string]int = make(map[string]int, 100)
	d["a"] = 1000
	d["b"] = 2000

	a = d
	fmt.Printf("%T, %v \n", a, a)
}

// 接受的参数为interface, 空接口传参
func describe(a interface{}) {
	fmt.Printf("%T, %v \n", a, a)
}

type interfaceStudent struct {
	Name string
	Age  int
}

func testInterface2() {
	str := "jarvis"
	describe(str)

	stu := interfaceStudent{
		Name: "jarvis",
		Age:  35,
	}
	describe(stu)
}

// 类型的断言
func assertion(a interface{}) {
	//通过ok 的语法，a.(T) 进行判断
	s, ok := a.(int)
	if !ok {
		fmt.Printf("a can't conver to int \n")
		return
	}
	fmt.Println(s)
}

func main() {

	testInterface()

	testInterface2()

	var a int = 100
	assertion(a)

	var b string = "hello"
	assertion(b)

}
