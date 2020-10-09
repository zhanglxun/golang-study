package main

import (
	list2 "container/list"
	"fmt"
)

/**
  基本定义和使用
*/

func testList1() {

	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)
	fmt.Printf("len: %v \n", list.Len())
	fmt.Printf("first:%v \n", list.Front().Value)
	fmt.Printf("second:%v \n", list.Front().Next().Value)

}

func main() {

	testList1()

}
