package main

import (
	"io"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() (league []Player) {
	f.database.Seek(0, 0)
	league, _ = NewLeague(f.database)
	return
}
func (f *FileSystemStore) GetPlayerScore(name string) (score int) {
	var wins int
	for _, player := range f.GetLeague() {

	}
	return 0
}
