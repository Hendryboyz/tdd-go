package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "henry")

	actual := buffer.String()
	expected := "Hello, henry"

	if actual != expected {
		t.Errorf("Got %q but expect %q", actual, expected)
	}
}
