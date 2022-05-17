package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there,I love %s", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	//fmt.Println("hello world")
}
