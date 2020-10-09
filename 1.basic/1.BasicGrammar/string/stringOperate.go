package main

import (
	"fmt"
	"golangStudy/1.basic/1.BasicGrammar/access"
)

//注意：只有首字母大写的，是共有的，可以在外部访问
func accessCall() {
	fmt.Printf("access is %s \n", access.Jerry)

	fmt.Printf("access is %s \n", access.Jarvis)

	access.StringPrint()

}

func main() {

	accessCall()
}
