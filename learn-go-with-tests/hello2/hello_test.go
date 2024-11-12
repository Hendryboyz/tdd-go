package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		actual := Hello("Chris", "")
		expect := "Hello, Chris"

		assertCorrectMessage(t, actual, expect)
	})

	t.Run("Say 'Hello, World' when an empty name is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expect := "Hello, World"

		assertCorrectMessage(t, actual, expect)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Henry", "Spanish")
		expect := "Hola, Henry"

		assertCorrectMessage(t, actual, expect)
	})

	t.Run("in French", func(t *testing.T) {
		actual := Hello("Henry", "French")
		expect := "Bonjour, Henry"

		assertCorrectMessage(t, actual, expect)
	})
}

func assertCorrectMessage(t testing.TB, actual, expect string) {
	// using Helper to avoid testing framework reporting the error inside test helper
	t.Helper()
	if actual != expect {
		t.Errorf("got %q expect %q", actual, expect)
	}
}
