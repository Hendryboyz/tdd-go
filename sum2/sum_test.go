package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		assertArrayError(t, actual, expected, numbers)
	})

	t.Run("Sum collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		assertArrayError(t, actual, expected, numbers)
	})
}

func assertArrayError(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %d want %d, %v", actual, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Sum 2 slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		assertSliceError(t, actual, expected)
	})

	t.Run("Sum 1 slice with lot of 1", func(t *testing.T) {
		actual := SumAll([]int{1, 1, 1, 1, 1, 1})
		expected := []int{6}

		assertSliceError(t, actual, expected)
	})
}

func assertSliceError(t testing.TB, actual, expected []int) {
	t.Helper()
	// only available in go 1.21
	// if !slices.Equal(actual, expected) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSumTail(t *testing.T) {

	checkSums := func(t testing.TB, actual, expected []int) {
		t.Helper()
		// only available in go 1.21
		// if !slices.Equal(actual, expected) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("Sum tail length 2 slice", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, actual, expected)
	})

	t.Run("Sum empty slices safely", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{1, 3, 9}, []int{2})
		expected := []int{0, 12, 0}

		checkSums(t, actual, expected)
	})
}
