package main

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	return
}
