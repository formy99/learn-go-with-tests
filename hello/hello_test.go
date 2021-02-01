package main

import "testing"

func TestHello(t *testing.T) {

	assetCorrentMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assetCorrentMessage(t, got, want)
	})

	t.Run("say hello, world when an empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assetCorrentMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assetCorrentMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Tony", "French")
		want := "Bejour, Tony"
		assetCorrentMessage(t, got, want)
	})

}
