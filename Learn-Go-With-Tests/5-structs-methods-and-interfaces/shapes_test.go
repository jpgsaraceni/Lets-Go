package main

import "testing"

var shapes = []struct {
	shape         Shape
	wantArea      float64
	wantPerimeter float64
}{
	{shape: Rectangle{width: 12, height: 6}, wantArea: 72.0, wantPerimeter: 36},
	{shape: Circle{radius: 10}, wantArea: 314.1592653589793, wantPerimeter: 62.83185307179586},
}

func TestArea(t *testing.T) {
	for _, tt := range shapes {
		got := tt.shape.Area()
		if got != tt.wantArea {
			t.Errorf("%#v got %g want %g", got, tt.wantArea)
		}
	}
}

func TestPerimeter(t *testing.T) {
	for _, tt := range shapes {
		got := tt.shape.Perimeter()
		if got != tt.wantPerimeter {
			t.Errorf("%#v got %g want %g", got, tt.wantPerimeter)
		}
	}
}
