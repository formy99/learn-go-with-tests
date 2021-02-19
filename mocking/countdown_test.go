package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &ConfigurableSleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountDownOperationSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("got %v ,want %v", spySleeperPrinter.Calls, want)
		}
	})

}
