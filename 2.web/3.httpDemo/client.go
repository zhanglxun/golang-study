package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

func main() {
	clinet := http.Client{}

	resp, err := clinet.Get("http://localhost:8080/golang")
	if err != nil {
		println("error")
		return
	}
	fmt.Println(resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Println(data)

	strA := "date"
	strB := "你好"
	fmt.Printf("长度 %d \n", len(strA))
	fmt.Printf("长度 %d \n", len(strB))
	fmt.Printf("长度 %d \n", utf8.RuneCountInString(strB))

}
