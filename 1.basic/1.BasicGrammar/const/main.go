package main

import "fmt"

func main() {

	//常量使用const 修饰，代表只读，不能修改
	const (
		a int = 100
		b
		c int = 200
		d
		x
	)
	fmt.Printf("a=%d b=%d c=%d d=%d x=%d \n", a, b, c, d, x)

	const (
		e = iota
		f
		g
	)

	//iota是golang语言的常量计数器,只能在常量的表达式中使用.
	const (
		a1 = 1 << iota
		a2 = 1 << iota
		a3 = 1 << iota
	)

	fmt.Printf("e=%d f=%d g=%d\n", e, f, g)

	fmt.Printf("a1=%d a2=%d a3=%d\n", a1, a2, a3)

}
