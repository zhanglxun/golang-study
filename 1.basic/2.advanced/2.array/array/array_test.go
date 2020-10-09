package array

import "testing"

func TestArray2(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want [2]int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("ok", func(t *testing.T) {
		got := Array2()
		want := [2]int{1, 2}
		assertCorrectMessage(t, got, want)
	})

}
