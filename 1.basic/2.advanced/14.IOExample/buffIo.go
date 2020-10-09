package main

import (
	"bufio"
	"fmt"
	"os"
)

func mytest() {
	fmt.Println("this is a good day")
}

func main() {

	var str string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Println("read from bufio:", str)
}
