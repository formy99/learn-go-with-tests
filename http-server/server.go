package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerSerer struct {
	store PlayerStore
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	playerScore := GetPlayerScore(player)
	fmt.Fprint(w, playerScore)

}

func GetPlayerScore(name string) string {
	switch name {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	default:
		return ""
	}

}
