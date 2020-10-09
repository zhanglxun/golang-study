package main

import (
	"fmt"
	"math"
)

/**
定义了一个结构体
*/
type Vertex struct {
	X, Y float64
}

/**
go 里面没有在其他语言那样，在class 里面写很多方法的模式，但是也实现了。
在普票函数前面，价格接收者，receiver 就 知道这个函数，属于哪个结构体struct 了
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(100)
	fmt.Println(v.Abs())
}
