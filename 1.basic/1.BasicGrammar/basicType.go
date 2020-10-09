package main

import "fmt"

const Pi = 3.1415926

func main() {
	v := 42

	j := v

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("type of v is %T\n", v)
	fmt.Printf("type of i is %T\n", i)
	fmt.Printf("type of f is %T\n", f)
	fmt.Printf("type of g is %T\n", g)
	fmt.Printf("type of j is %T\n", j)

	//常量不能用 := 语法声明，需要用const 进行声明，常亮不能修改
	const World = "世界"
	const name = "zhanglixun"

	fmt.Println("Hello", World)
	fmt.Println("hello", name)
	fmt.Println("Happy", Pi, "Day")

	const (
		a int = 100
		b
		c int = 200
		d
	)
	fmt.Printf("a= %d,b=%d,c=%d,d=%d\n", a, b, c, d)

	const (
		e = iota
		k
		n
	)
	const (
		a1 = 1 << iota
		a2 = 1 << iota
		a3 = 1 << iota
	)
	fmt.Printf("e= %d,f= %d,g= %d\n", e, k, n)
	fmt.Printf("a1=%d,a2=%d,a3=%d\n", a1, a2, a3)

	m1 := [3]int{1, 2, 3}
	fmt.Println(m1)
	fmt.Printf("type of m1 is %T\n", m1)

	m2 := []int{1, 2}
	m2 = append(m2, 3)
	fmt.Println(m2)
	fmt.Printf("type of m1 is %T\n", m2)

}
