package main

/**
取的是 计数的循环下标，0，1，2
*/
func testFor1() {

	arr := []int{11, 12, 13}
	for i, _ := range arr {
		println(i)
	}
}

/**
取的是 循环下标 的值 ，14，15，16
*/
func testFor2() {

	arr := []int{14, 15, 16}
	for _, v := range arr {
		println(v)
	}
}

func main() {

	testFor1()

	testFor2()
}
