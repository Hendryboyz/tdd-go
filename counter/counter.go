package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord   = "Go!"
	startNumber = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := startNumber; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	CountDown(os.Stdout, sleeper)
}
