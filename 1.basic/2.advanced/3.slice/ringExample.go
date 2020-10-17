package main

import "fmt"

func taseMake1() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	//append是追加到最后面
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	//满了之后，写第10 个，没问题，到11的时候，要扩容10
	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}

	//观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

}

func testMake12() {
	var a []int
	a = make([]int, 5, 10)
	//a[5] = 100
	a = append(a, 10)
	fmt.Printf("a=%v\n", a)

	b := make([]int, 0, 10)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))
	b = append(b, 100)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))
}

func main() {
	//taseMake1()
	testMake12()
}
