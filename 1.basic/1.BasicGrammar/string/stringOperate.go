package main

import (
	"fmt"
	"golangStudy/1.basic/1.BasicGrammar/access"
	"unicode/utf8"
)

//注意：只有首字母大写的，是共有的，可以在外部访问
func accessCall() {
	fmt.Printf("access is %s \n", access.Jerry)

	fmt.Printf("access is %s \n", access.Jarvis)

	access.StringPrint()

}

func stringLen() {
	var name = "你好"
	println(len(name))
	//计算字符个数，要用这个。
	println(utf8.RuneCountInString(name))
}

func main() {

	accessCall()
	stringLen()
}
