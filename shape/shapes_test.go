package shape

import "testing"

func TestPreimeter(t *testing.T) {
	t.Run("Given Rectangle", func(t *testing.T) {
		rect := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		actual := Preimeter(rect)
		expected := 40.0
		assertFloat64(t, expected, actual)
	})
}

func assertFloat64(t testing.TB, expected, actual float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expect %.2f, but get %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("Area of %#v expect %g, but get %g", shape, expected, actual)
		}
	}

	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, expectedArea: 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			checkArea(t, testCase.shape, testCase.expectedArea)
		})
	}
}
