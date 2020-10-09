package main

import "fmt"

// 变量相关操作

func main() {

	var name = "jarvis!"

	fmt.Println(name)

	var (
		a int
		b bool
		c string
		d float32
	)

	const f int = 100
	const Pi float32 = 3.1415

	fmt.Printf("a=%d b=%t c=%s d=%f f=%d\n", a, b, c, d, f)

	a = 10
	b = true
	c = "hello"
	d = 10.8

	fmt.Printf("a=%d b=%t c=%s d=%f\n", a, b, c, d)

	fmt.Println(f, Pi)

	//引用类型的赋值
	ch := make(chan int)
	fmt.Println(ch)

}
