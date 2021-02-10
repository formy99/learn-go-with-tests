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
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Circle{10.00}, want: 314.1592653589793},
		{shape: Rectangle{10.00, 10.00}, want: 100.00},
		{shape: Triangle{3.00, 4.00}, want: 6.00},
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %f, want %f", tt.shape, got, tt.want)
		}
	}
}
