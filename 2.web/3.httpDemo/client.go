package main

import (
	"fmt"
	"net/http"
)

func main() {
	clinet := http.Client{}

	resp, err := clinet.Get("http://localhost:8080/golang")
	if err != nil {
		println("error")
		return
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.Body)

}
