package generic_test

import (
	"generic"
	"testing"
)

func AssertTrue(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Errorf("got %v, want true", actual)
	}
}

func AssertFalse(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Errorf("got %v, want false", actual)
	}
}

func TestStack(t *testing.T) {
	t.Run("test integer stack", func(t *testing.T) {
		// stack := new(generic.StackOfInt)
		stack := generic.StackOfInt{}
		AssertTrue(t, stack.IsEmpty())

		stack.Push(123)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(456)
		val, _ := stack.Pop()
		AssertEqual(t, 456, val)

		val, _ = stack.Pop()
		AssertEqual(t, 123, val)
		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("test string stack", func(t *testing.T) {
		stack := new(generic.StackOfString)

		AssertTrue(t, stack.IsEmpty())

		stack.Push("hello")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("world")
		val, _ := stack.Pop()
		AssertEqual(t, "world", val)

		val, _ = stack.Pop()
		AssertEqual(t, "hello", val)
		AssertTrue(t, stack.IsEmpty())
	})
}
