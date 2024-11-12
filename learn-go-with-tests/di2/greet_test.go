package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Henry")
	got := buffer.String()
	expected := "Hello, Henry"

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}
