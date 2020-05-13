package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

const jsonContentType = "application/json"

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
	router.Handle("/game", http.HandlerFunc(s.gameHandler))
	router.Handle("/ws", http.HandlerFunc(s.webSocketHandler))
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
	league := s.GetLeagueTable()
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(league)
}

func (s *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("game.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("problema loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (s *PlayerServer) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _ := upgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	s.store.RecordWin(string(winnerMsg))
}

func (s *PlayerServer) GetLeagueTable() []Player {
	league := s.store.GetLeague()
	sort.Slice(league, func(i, j int) bool {
		return league[i].Wins > league[j].Wins
	})
	return league

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
