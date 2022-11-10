package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum 5 numbers", func(t *testing.T) {
		// numbers := [5]int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		assertSumSlice(t, expected, actual, numbers)
	})

	t.Run("Sum 6 numbers", func(t *testing.T) {
		numbers := []int{3, 6, 9, 12, 15, 18}
		actual := Sum(numbers)
		expected := 63
		assertSumSlice(t, expected, actual, numbers)
	})
}

func assertSumSlice(t testing.TB, expected int, actual int, slice []int) {
	if actual != expected {
		t.Errorf("Expected %d but actual %d, %v", expected, actual, slice)
	}
}
