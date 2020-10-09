package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
	//println(got, want)
}
