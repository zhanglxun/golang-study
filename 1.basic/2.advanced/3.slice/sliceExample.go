package main

import (
	"fmt"
	"sort"
)

/*
 切面是对数组的封装，是可变灵活的
*/

/**
  切片需要赋值才能操作，
  切面和数组的区别，在于数据的长度

*/
func testSlice0() {
	var a []int
	if a == nil {
		fmt.Printf("a is nil \n")
	} else {
		fmt.Printf("a is %v \n", a)
	}
}

//切片从数组中拿值，b[1:4]从下标为，取到下标4 前面一个
func testSlice1() {
	var b [6]int = [6]int{1, 2, 3, 4, 5, 6}
	c := b[1:4]
	fmt.Printf("slice c:%v \n", c)

	var d []int = []int{9, 8, 7}
	fmt.Printf("slice d:%v \n", d)

	e := b[1:]
	f := b[:4]
	g := b[:]

	fmt.Printf("%d,%d,%d \n", e, f, g)
}

//对切片取值后，值进行操作
func testSlice2() {
	a := [...]int{1, 2, 3, 4, 5, 7, 8, 9, 11}
	fmt.Printf("array a: %d,type is %T \n", a, a)

	b := a[2:5]
	fmt.Printf("array b: %d,type is %T \n", b, b)

	b[0] = b[0] + 10
	b[1] = b[1] + 20
	b[2] = b[2] + 30
	fmt.Printf("b=%d \n", b)
}

func testLen() {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

}

// make 对切片
func testMake() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	a[1] = 12
	fmt.Printf("a=%v \n", a)

}

// 切片的扩容,当容量 cap 满之后，扩展一倍
func testMake2() {
	var a []int
	a = make([]int, 1, 10)
	a[0] = 10
	fmt.Printf("a=%v ,addr=%p,len=%d,cap=%d \n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v ,addr=%p,len=%d,cap=%d \n", a, a, len(a), cap(a))

	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("a=%v ,addr=%p,len=%d,cap=%d \n", a, a, len(a), cap(a))
	}
}

// 切片后的容量,切片的容量，是数组是起始位置，到结尾的大小
func testSliceCut() {
	apple := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	b := apple[1:3]
	fmt.Printf("b=%v ,addr=%p,len=%d,cap=%d \n", b, b, len(b), cap(b))
	c := b[:cap(b)]
	fmt.Printf("c=%v ,addr=%p,len=%d,cap=%d \n", c, c, len(c), cap(c))

}

//切片展开append,使用 ... 展开元素,
/*
【重点：切片是引用类型】
*/
func testSliceAppend() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8}
	a = append(a, 21, 22, 23)
	fmt.Printf("a=%d \n", a)
	a = append(a, b...)
	fmt.Printf("a after append is :%d \n", a)

}

// 使用切片传参，数组数据累加
func sumArraySlice(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

/**
调用数组累加
*/
func testSumSlice() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sumArraySlice(a[:])
	fmt.Printf("sum =%d \n", sum)
}

// 切片的拷贝 ,被拷贝的对象，不会扩容，会减少
func testSliceCopy() {
	a := []string{"apple", "orange"}
	b := []string{"cat", "dog", "chicken"}
	copy(a, b)
	fmt.Printf("a=%s \n", a)
	fmt.Printf("b=%s \n", b)
}

// make 和new 的区别
/*
 make 为内建的类型 slice，map, channel 分配内存
 new 用于各种类型的内存分配，new 返回是一个指针
*/

//数组进行排序
func sortArray() {
	var a = [7]int{3, 4, 6, 2, 5, 9, 8}
	sort.Ints(a[:])
	fmt.Printf("a=%d \n", a)

	var b = [5]string{"a", "c", "e", "d", "b"}
	sort.Strings(b[:])
	fmt.Printf("b=%s \n", b)
}

func main() {

	testSlice0()

	testSlice1()

	testSlice2()

	testLen()

	testMake()

	testMake2()

	testSliceCut()

	testSliceAppend()

	testSumSlice()

	testSliceCopy()

	sortArray()

}
