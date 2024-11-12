package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountDownOperations struct {
	Operations []string
	Calls      int
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls++
	s.Operations = append(s.Operations, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Operations = append(s.Operations, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountDownOperations{}
		CountDown(buffer, sleeper)

		got := buffer.String()
		expected := `3
2
1
Go!
`

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}

		if sleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", sleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		sleepPrinter := &SpyCountDownOperations{}
		CountDown(sleepPrinter, sleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, sleepPrinter.Operations) {
			t.Errorf("wanted calls %v got %v", want, sleepPrinter.Operations)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
