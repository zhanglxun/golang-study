package main

import (
	"errors"
	"fmt"
)

func main() {

	var e error
	e = errors.New("this is a test error")

	fmt.Println(e.Error()) // 输出：This is a test error

	// 使用 fmt.Errorf() 定制错误信息
	err := fmt.Errorf("This is another error")
	fmt.Println(err) // 输出：This is another test error
}
