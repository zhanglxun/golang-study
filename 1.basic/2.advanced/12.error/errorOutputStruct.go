package main

import "fmt"

type errorStringStruct struct {
	s string
}

func (e errorStringStruct) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorStringStruct{text}
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		fmt.Println("Error：", ErrType)
	} else {
		fmt.Println("Error:", nil)
	}
}
