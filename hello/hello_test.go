package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("When name is Henry Then greeting to Henry", func(t *testing.T) {
		actual := Hello("Henry", "")
		expected := "Hello, Henry"
		assertCorrectString(t, expected, actual)
	})
	t.Run("When name is empty Then name should be world", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectString(t, expected, actual)
	})
	t.Run("When in Spanish", func(t *testing.T) {
		actual := Hello("Henry", "Spanish")
		expected := "Hola, Henry"
		assertCorrectString(t, expected, actual)
	})
	t.Run("When in French", func(t *testing.T) {
		actual := Hello("Henry", "French")
		expected := "Bonjour, Henry"
		assertCorrectString(t, expected, actual)
	})
}

func assertCorrectString(t testing.TB, expect, actual string) {
	t.Helper() // announce helper func to prevent test from addressing the wrong line to fix
	if actual != expect {
		t.Errorf("Got %s want %s", actual, expect)
	}
}
