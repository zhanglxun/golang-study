package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
