package structs

import (
	"testing"
)

func assertCorrentMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f ,want %f", got, want)
	}
}
func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.00, 10.00}
	got := Perimeter(r)
	want := 40.00
	assertCorrentMessage(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		r := Rectangle{10.00, 10.00}
		checkArea(t, r, 100.00)
	})

	t.Run("Circle Area", func(t *testing.T) {
		c := Circle{10.00}
		checkArea(t, c, 314.1592653589793)
	})

}
