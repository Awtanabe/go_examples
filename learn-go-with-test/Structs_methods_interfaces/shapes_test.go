package main

import "testing"

// 初期テスト
// func TestPerimeter(t *testing.T) {
// 	r := Rectangle{10.0,10.0}
// 	got := r.Area()
// 	want := 20.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// テスト基本
// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12,6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("Circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// テーブルテスト
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape // 値
		want float64 // 欲しい結果
	}{
		{ Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
	   got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got , tt.want)
		}
	}
}