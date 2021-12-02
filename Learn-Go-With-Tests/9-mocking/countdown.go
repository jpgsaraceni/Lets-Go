package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(500 * time.Millisecond)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprintln(out, finalWord)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}
