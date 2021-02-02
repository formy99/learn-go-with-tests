package integers

import (
	"fmt"
	"testing"
)

// Add takes two integers and returns the sum of them.
func TestAdder(t *testing.T) {
	assertCorrentMessage := func(t testing.TB, got, want int) {
		if got != want {
			t.Errorf("got %d ,want %d", got, want)
		}
	}

	t.Run("add", func(t *testing.T) {
		got := Adder(1, 1)
		want := 2
		assertCorrentMessage(t, got, want)
	})
}

func ExampleAdder() {
	sum := Adder(1, 2)
	fmt.Println(sum)
	// Output: 3
}
