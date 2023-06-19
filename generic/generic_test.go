package generic_test

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integer", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual[string](t, "hello", "hello")
		AssertNotEqual[string](t, "hello", "greeting")
	})
}

func AssertEqual[T comparable](t *testing.T, expect, actual T) {
	t.Helper()
	if expect != actual {
		t.Errorf("got %+v, want %+v\n", actual, expect)
	}
}

func AssertNotEqual[T comparable](t *testing.T, expect, actual T) {
	t.Helper()
	if expect == actual {
		t.Errorf("didn't want %+v\n", actual)
	}
}
