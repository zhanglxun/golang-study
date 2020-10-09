package main

import "fmt"

//golan 中的切片

func main() {

	var slice1 []int = make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for i, value := range slice1 {

		fmt.Printf("slice at %d is : %d \n", i, value)
	}

	var slice2 []int = make([]int, 5)

	slice2[0] = 3
	slice2[1] = 2
	slice2[2] = 9
	slice2[2] = 10

	for j, value := range slice2 {
		fmt.Printf("slice2 at  #%d is :%d \n", j, value)
	}
}
