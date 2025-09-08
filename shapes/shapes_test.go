package shapes

import (
	"testing"
)

func compareResults(t testing.TB, expected, actual float64) {
	t.Helper()
	if expected != actual {
		t.Errorf("got %g expected %g ", expected, actual)
	}

}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	expected := 40.0
	actual := Perimeter(rectangle)
	compareResults(t, expected, actual)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, expected: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("%#v got %g want %g", tt.shape, actual, tt.expected)
			}
		})
	}
}
