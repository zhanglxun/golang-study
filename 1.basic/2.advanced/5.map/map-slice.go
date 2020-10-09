package main

import "fmt"

func test_map(m map[int]int) {
	m[0] = 1
}

func test_slice(a []int) {
	a = append(a, 1)
}

func main() {
	m := make(map[int]int)
	a := make([]int, 0, 10)

	test_map(m)
	test_slice(a)

	fmt.Println(len(m)) // 1
	fmt.Println(len(a)) // 0
}
