package main

import "fmt"

func iotaTest() {

	const a1 = iota
	fmt.Println("a1=", a1)

	// 每次 const 出现时，都会让 iota 初始化为0.如果在一个 const 里面定义，则不会
	const b1 = 1 << iota
	fmt.Println("b1=", b1)

	const c1 = 1 << iota
	fmt.Println("c1=", c1)
}

// 此定义方式，为掩码的表达式
const (
	a = 1 << iota
	b = 1 << iota
	c = 1 << iota
)

//自增常量，包含一个自定义的枚举类型，允许你依靠编译器完成自增设置
const (
	e = iota
	f
	g
	h
)

//常量有跳过的情况，则重新计数

const (
	A = iota
	B = iota
	C = iota
	D = 8
	E = 8
	F = iota
	G = iota
)
const (
	A1 = iota
	A2
)

func printConst() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)

	fmt.Println("A1A2")
	fmt.Println(A1)
	fmt.Println(A2)
}

func main() {

	iotaTest()

	fmt.Printf("a=%d b=%d c=%d \n", a, b, c)

	//自增的一个定义方式
	fmt.Printf("e=%d f=%d g=%d h=%d \n", e, f, g, h)

	//有间隔的情况
	printConst()

}
