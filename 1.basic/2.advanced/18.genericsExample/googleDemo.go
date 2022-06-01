package main

import "fmt"

func SumInt32(x, y int32) int32 {
	return x + y
}
func Float64(x, y float64) float64 {
	return x + y
}

func Add[T int64 | float64](x, y T) T {
	return x + y
}

func main() {

	fmt.Println(SumInt32(30, 40))
	fmt.Println(Float64(30.79, 43.77))
	fmt.Println(Add(31.79, 41.77))

}
