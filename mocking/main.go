package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord  = "Go!"
	countStart = 3
	write      = "write"
	sleep      = "sleep"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
}

type CountDownOperationSpy struct {
	Calls []string
}

func (s *CountDownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)

}
