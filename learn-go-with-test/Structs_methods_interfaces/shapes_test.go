package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	testfunc := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}	
	}

	got := Perimeter(Rectangle{ 10.0, 10.0})
	want := 40.0

	testfunc(t, got, want)
}

// func TestArea(t *testing.T) {
// 	checkArea := func (t *testing.T, got, want float64) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("test rectngle area", func (t *testing.T) {
// 		got := Rectangle{10.0, 15.0}.Area()
// 		want := 150.0
// 		checkArea(t, got, want)
// 	})

// 	t.Run("testing circle area", func (t *testing.T){
// 		got := Circle{10.0}.Area()
// 		want := 314.1592653589793
// 		checkArea(t, got, want)
// 	})
// }

// テーブル駆動開発

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape // インターフェース
		hasArea float64
}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
}

for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()
				if got != tt.hasArea {
						t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
				}
		})

}

}