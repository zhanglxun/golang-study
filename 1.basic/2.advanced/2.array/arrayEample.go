package main

import (
	"fmt"
	"math/rand"
	"time"
)

//定义数组的第一种写法：变量+ 类型
func testArray() {
	var a [5]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [5]int = [5]int{10, 20, 30}

	fmt.Println(b)

	c := [3]int{1, 2, 3}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	e := [5]int{3: 100, 4: 300}
	fmt.Println(e)
}

//二维数组的输出
func testArray2() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	//数组，常规循环
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//数组按 range 范围进行循环取值 不应下标，可以用下划线 _ 把index 下标忽略掉
	for index, value := range a {
		fmt.Printf("a[%d]=%d,", index, value)
	}
	fmt.Println()
}

//二维数组定义，输出
func testArray3() {
	var a [2][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40

	fmt.Println(a)

	for i, value := range a {
		for j, value2 := range value {
			fmt.Printf("[%d,%d] =%d ,", i, j, value2)
		}
		fmt.Println()
	}

	b := [3][4]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "duck"},
	}
	fmt.Println(b)

}

//数组 是值类型，赋值，修改，原值不变
func testArray4() {
	a := [3]int{10, 20, 30}
	b := a
	b[0] = 1000

	fmt.Printf("a=%d,b=%d \n", a, b)
}

//数组传参,是把值 复制过去了，并没有改变原值
func modify(b [3]int) {
	b[0] = 1000
}

func testArray5() {
	a := [3]int{10, 20, 30}

	modify(a)
	fmt.Println(a)

}

//练习，求数组的 元素的和
func sumArray(a [10]int) (sum int) {
	sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return
}
func testArray6() {
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(100)
	}
	sum := sumArray(b)
	fmt.Printf("数组:%d 之和为:%d \n", b, sum)
}

//这个是 某个 google 的工程师，做人工智能出过的题目
// 找出数组中，给定值的下标
func testArray7(a [10]int, target int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i]+a[j] == target {

				fmt.Printf("(%d,%d)", a[i], a[j])
			}
		}
	}

}

func main() {

	testArray()
	//数组的使用
	testArray2()

	testArray3()

	testArray4()

	testArray5()

	testArray6()

	var test [10]int = [10]int{1, 3, 5, 8, 7, 2, 7, 6, 0, 9}
	testArray7(test, 8)
}
