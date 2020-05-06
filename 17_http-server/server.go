package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := new(PlayerServer)
	s.store = store
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(s.playerHandler))
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	s.Handler = router

	return s
}

func (s *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		s.ProcessWin(w, player)
	case http.MethodGet:
		s.ShowScores(w, player)
	}
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := s.GetLeagueTable()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(leagueTable)
}

func (s *PlayerServer) GetLeagueTable() []Player {
	return s.store.GetLeague()

}

func (s *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {
	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) ShowScores(w http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}
