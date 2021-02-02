package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	assertCorrentMessage := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	got := Repeater("a", 6)
	want := "aaaaaa"
	assertCorrentMessage(t, got, want)
}

func ExampleRepeater() {
	repstr := Repeater("b", 6)
	fmt.Println(repstr)
	// Output: bbbbbb
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("b", 10)
	}
}
