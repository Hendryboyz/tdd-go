package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("When name is Henry Then greeting to Henry", func(t *testing.T) {
		result := Hello("Henry")
		expected := "Hello, Henry"
		assertCorrectString(t, result, expected)
	})
	t.Run("When name is empty Then name should be world", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"
		assertCorrectString(t, expected, result)
	})
}

func assertCorrectString(t testing.TB, expect, actual string) {
	t.Helper() // announce helper func to prevent test from addressing the wrong line to fix
	if actual != expect {
		t.Errorf("Got %s want %s", actual, expect)
	}
}
