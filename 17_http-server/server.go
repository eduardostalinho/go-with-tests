package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
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

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		s.ProcessWin(w, player)
	case http.MethodGet:
		s.ShowScores(w, player)
	}
}
