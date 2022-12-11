package main

import (
	"fmt"
	"io"
	"os"
)

const countdownStart = 3
const finalWord = "Go!"

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
