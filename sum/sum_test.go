package sum

import (
	"reflect"
	"testing"
)

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
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %d but actual %d, %v", expected, actual, slice)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("sum of 2 slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but actual %v", expected, actual)
		}
	})

	t.Run("sum of 1 slices", func(t *testing.T) {
		actual := SumAll([]int{1, 1, 1})
		expected := []int{3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, but actual %v", expected, actual)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, actual []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expect %v, actual %v", expected, actual)
		}
	}
	t.Run("Sum 2 slices tails", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, actual, expected)
	})

	t.Run("Sum empty slice safely", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{2, 4, 6}, []int{1})
		expected := []int{0, 10, 0}
		checkSums(t, actual, expected)
	})
}
