package main

import (
	"fmt"
)

func SumInt32(x, y int32) int32 {
	return x + y
}
func Float64(x, y float64) float64 {
	return x + y
}

//使用泛型进行输出
func Add[T int64 | float64](x, y T) T {
	return x + y
}

//使用泛型进行输出
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println("")
}

type Addable interface {
	int | int8 | int32 | float32 | float64 |
		complex64 | complex128 | string
}

//type Addable2 interface {type int,interface{} }

func addGenerics[T Addable](a, b T) T {
	return a + b
}

//type Addable interface{
//	type int, int8, int16, int32, int64,
//	uint, uint8, uint16, uint32, uint64, uintptr,
//	float32, float64, complex64, complex128,    string
//}

func main() {

	fmt.Println(SumInt32(30, 40))
	fmt.Println(Float64(30.79, 43.77))
	fmt.Println(Add(31.79, 41.77))

	printSlice[int]([]int{1, 2, 3, 4, 5})
	printSlice[float64]([]float64{1.01, 2.02, 3.03, 4.04, 5.05})
	printSlice([]string{"Hello", "World"})
	printSlice[int64]([]int64{5, 4, 3, 2, 1})

	fmt.Println(addGenerics(3.33, 6.98))
	fmt.Println(addGenerics(30, 40))
	fmt.Println(addGenerics("jarvis", "hello"))
}
