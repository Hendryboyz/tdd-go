package shape

import "testing"

func TestPreimeter(t *testing.T) {
	t.Run("Given width and height", func(t *testing.T) {
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
	if actual != expected {
		t.Errorf("Expect %.2f, but get %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	t.Run("Given width and height", func(t *testing.T) {
		rect := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		actual := Area(rect)
		expected := 72.0
		assertFloat64(t, expected, actual)
	})
}
