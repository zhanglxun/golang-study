package main

import "os"
import "fmt"

func main() {

	fmt.Println(os.Args)
	//fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		fmt.Println("Hello world", os.Args[0])
	}
	//os.Exit(0)
}
