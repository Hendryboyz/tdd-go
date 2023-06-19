package generic_test

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integer", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "greeting")
	})
}

func AssertEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()
	if expect != actual {
		t.Errorf("got %d, want %d\n", actual, expect)
	}
}

func AssertNotEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()
	if expect == actual {
		t.Errorf("didn't want %d\n", actual)
	}
}
