package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "a", Age: 24},
		{Name: "b", Age: 23},
		{Name: "c", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}

}

func pase_student_right() {

	m := make(map[string]*student)
	stus := []student{
		{Name: "d", Age: 24},
		{Name: "e", Age: 23},
		{Name: "f", Age: 22},
	}

	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}

}

func Test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	//for range 循环的时候会创建每个元素的副本，而不是元素的引用
	//所以 m [key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，
	//因为最后 val 被赋值为 3，所有输出都是 3.
	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Printf("key: %d, value: %d \n", k, *v)
	}
}

func main() {

	pase_student()

	pase_student_right()

	Test2()
}
