package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tom")

	got := buffer.String()
	want := "Hello, Tom"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
