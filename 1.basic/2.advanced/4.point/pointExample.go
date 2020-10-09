package main

import "fmt"

/*
 1、变量和内存地址
 2、指针类型
 3、值拷贝和引用拷贝
*/
//内存地址 ，使用& 符号可获取地址  * 用来解析地址
func testPoint1() {
	var a int32 = 100
	fmt.Printf("a=%d,address=%v \n", a, &a)

	var b string = "jarvis"
	fmt.Printf("b=%s,address=%v \n", b, &b)

}

//指针类型，或叫引用类型,
func testPoint2() {
	b := 253
	var a *int = &b

	fmt.Printf("2a=%d,a address=%p \n", *a, &a)

	//指针 初始化不赋值，为nil
	var c *int
	fmt.Printf("2c = %v ,c addr is %p \n", c, &c)
}

func testPoint3() {
	var a int = 200
	var b *int = &a

	fmt.Printf("3b address 's value is :%d \n", *b)
	*b = 1000
	*b++
	fmt.Printf("3b address 's value is %d \n", *b)
	fmt.Printf("3a =%d \n", a)
}

/**
指针传参数，引用类型
*/
func change(val *int) {
	*val = 55 /* 声明指针变量 */
}

func changeValue() {
	a := 66
	fmt.Printf("a=%d \n", a)
	b := &a
	change(b)
	fmt.Printf("b=%d \n", a)
}

/**
指针的数组传参,使用 * 号，*号用于指定变量是作为一个指针
*/
func modify_address(val *[3]int) {
	(*val)[0] = 100
}

func testPoint4() {
	var b [3]int = [3]int{3, 5, 7}
	modify_address(&b)
	fmt.Printf("after change b=%d\n", b)
}

// 指针 切片传参,切片本身就是 引用类型，不用加* ，和上面的数组不同
func modifySlice(a []int) {
	a[0] = 90
}
func testPoint5() {
	a := [3]int{91, 92, 93}
	fmt.Printf("array is %d \n", a)
	modifySlice(a[:])
	fmt.Printf("slice param is %d \n", a)
}

//指针的 make 和 new 的区别
/*
	1,make 只能用于 slice,map,channel,make(T, args) 返回的是初始化之后的 T 类型的值，
	这个新值并不是 T 类型的零值,也不是指针 *T，是经过初始化之后的 T 的引用.
	2,new(T) 返回 T 的指针 *T,并指向 T 的零值。
*/
func testPoint6() {
	var a *int = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	var b *[]int = new([]int)
	fmt.Printf("*b=%v \n", *b)

	//使用make 初始化一个 slice 切片，初始值都是 0
	*b = make([]int, 5, 100)
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b=%v \n", *b)

}

/**
这是指针的一个例子，a是 int 类型 b 和c 都是指针类型，c 的值变化后，实际是直接给地址上的值做了修改
*/
func testPoint7() {
	var a int = 10
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("*c=%d,*b=%d,a=%d\n", *c, *b, a)
}

func main() {
	testPoint1()

	testPoint2()

	testPoint3()

	changeValue()

	testPoint4()

	testPoint5()

	testPoint6()

	testPoint7()

}
