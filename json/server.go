package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Handler.ServeHTTP(w, r)

}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("context-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())

}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Bob", 12},
	}
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

type StubPlayStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayStore) GetLeague() []Player {
	return s.league
}

type Player struct {
	Name string
	Wins int
}
