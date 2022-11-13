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
			t.Errorf("Expect %g, but get %g", expected, actual)
		}
	}
	t.Run("Given Rectangle", func(t *testing.T) {
		rect := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		expected := 72.0
		checkArea(t, rect, expected)
	})

	t.Run("Given Circle", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}
