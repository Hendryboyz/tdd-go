package geometry

import "testing"

func TestCalculatePerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.9}
	got := CalculatePerimeter(rectangle)
	expected := 41.8

	assertFloatNumbers(t, got, expected)
}

func assertFloatNumbers(t testing.TB, got, expected float64) {
	if got != expected {
		t.Errorf("got %g want %g", got, expected)
	}
}

func TestCalculateArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%#v got %g want %g", shape, got, hasArea)
		}
	}

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.9}, expected: 109.0},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.expected)
		})
	}

}
