package structure

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{ 10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := []struct{
		name string
		shape Shape
		want float64
	} {
		{ "Rectabgle", Rectangle{ 10.0, 10.0}, 100.0 },
		{ "Circle", Circle{ 10.0 }, 314.1592653589793 },
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}