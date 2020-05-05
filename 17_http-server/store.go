package main

type PlayerStore interface {
	GetPlayerScore(name string) int
}
