package shape

import "testing"

func TestPreimeter(t *testing.T) {
	t.Run("Given width and height", func(t *testing.T) {
		actual := Preimeter(10.0, 10.0)
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
		actual := Area(12.0, 6.0)
		expected := 72.0
		assertFloat64(t, expected, actual)
	})
}
