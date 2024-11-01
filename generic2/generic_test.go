package generic2

import "testing"

func TestAssertFunction(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual(t, "henry", "henry")
		AssertNotEqual(t, "henry", "chou")
	})

	// AssertEqual(t, 1, "1")
	// GenericFoo[string](1, "1") invalid
	// InterfaceFoo(1, "1") valid
}

func GenericFoo[T any](x, y T) {}

func InterfaceFoo(x, y any) {}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}
