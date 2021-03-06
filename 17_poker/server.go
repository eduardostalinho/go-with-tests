package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const jsonContentType = "application/json"
const htmlTemplatePath = "game.html"

type PlayerServer struct {
	http.Handler
	store    PlayerStore
	template *template.Template
	game     Game
}

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	tmpl, err := template.ParseFiles(htmlTemplatePath)

	if err != nil {
		return nil, fmt.Errorf("problema loading template %s %v", htmlTemplatePath, err)
	}

	s := new(PlayerServer)
	s.game = game
	s.template = tmpl
	s.store = store
	s.addHandlers()

	return s, nil
}

func (s *PlayerServer) addHandlers() {
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(s.playerHandler))
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/game", http.HandlerFunc(s.gameHandler))
	router.Handle("/ws", http.HandlerFunc(s.webSocketHandler))
	s.Handler = router
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
	s.template.Execute(w, nil)
}

func (s *PlayerServer) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	s.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()
	s.game.Finish(winner)
}

func (s *PlayerServer) GetLeagueTable() League {
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
