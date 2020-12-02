package main

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {

	if ErrNamedType == New("EOF") {
		fmt.Println("Named type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("struct type Error")
	}
}
