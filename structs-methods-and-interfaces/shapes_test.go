package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 20.00}
	got := Perimeter(rectangle)
	want := 60.00
	if got != want {
		t.Errorf("want %f ,get %f", want, got)
	}

}

func TestArea(t *testing.T) {
	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("want %f ,get %f", want, got)
	//	}
	//
	//}
	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{10.00, 20.00}
	//	want := 200.00
	//	checkArea(t, rectangle, want)
	//})
	//
	//t.Run("Circle", func(t *testing.T) {
	//	circle := Circle{10.00}
	//	want := 314.16
	//	checkArea(t, circle, want)
	//})
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 12.0},
		{shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{Width: 3, Height: 4}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v want %f ,get %f", tt.shape, tt.hasArea, got)
		}
	}
}
