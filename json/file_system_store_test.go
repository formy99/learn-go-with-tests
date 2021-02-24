package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
                 {"Name": "Cleo", "Wins": 10},
		         {"Name": "Chris", "Wins": 33}]`)
		store := FileSystemStore{database}
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
	})

	t.Run("getPlayerScore", func(t *testing.T) {
		database := strings.NewReader(`[
                 {"Name": "Cleo", "Wins": 10},
		         {"Name": "Chris", "Wins": 33}]`)
		store := FileSystemStore{database}
		got := store.GetPlayerScore("Chris")
		want := 33
		if got != want {
			t.Errorf("got %d ,want %d", got, want)
		}
	})
}
