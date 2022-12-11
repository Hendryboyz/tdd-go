package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	// action
	CountDown(buffer, sleeper)

	// assert
	actual := buffer.String()
	expected := `3
2
1
Go!
`

	if actual != expected {
		t.Errorf("Got %q, expected %q", actual, expected)
	}

	if sleeper.Calls != 3 {
		t.Errorf("Not enough calls to sleeper, epected 3 but actual %q", sleeper.Calls)
	}
}
