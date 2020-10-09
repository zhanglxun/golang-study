package main

import "fmt"

const englishHelloPrefix = "hello,"

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
