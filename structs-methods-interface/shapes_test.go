package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	// checkArea := func(t testing.TB, shape Shape, want float64) {

	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}

	// }

	// t.Run("Rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	got := Area(rectangle)
	// 	want := 72.0

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("Circles", func(t *testing.T) {
	// 	circle := Circles{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circles{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
