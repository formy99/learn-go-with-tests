package main

import "fmt"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, nil}
}

type InMemoryPlayerStore struct {
	store  map[string]int
	league []Player
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++

	fmt.Println(i.store)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league

}
