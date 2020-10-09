package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w, "%v\n", r.Form)
	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	//fmt.Fprintf(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

	}
	fmt.Printf("Hello World!")

}

func main() {

	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
		return
	}
}
