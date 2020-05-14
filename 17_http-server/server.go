package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

const jsonContentType = "application/json"
const htmlTemplatePath = "game.html"

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     IGame
}

type playerServerWS struct {
	*websocket.Conn
}

type Player struct {
	Name string
	Wins int
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("could not upgrade connection, %v\n", err)
	}

	return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("could not read message from websocket %v\n", err)
	}
	return string(msg)

}

func NewPlayerServer(store PlayerStore, game IGame) (*PlayerServer, error) {
	tmpl, err := template.ParseFiles(htmlTemplatePath)

	if err != nil {
		return nil, fmt.Errorf("problema loading template %s %v", htmlTemplatePath, err)
	}

	s := new(PlayerServer)
	s.game = game
	s.template = tmpl
	s.store = store
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(s.playerHandler))
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/game", http.HandlerFunc(s.gameHandler))
	router.Handle("/ws", http.HandlerFunc(s.webSocketHandler))
	s.Handler = router

	return s, nil
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

	s.game.Start(numberOfPlayers, ioutil.Discard)
	winner := ws.WaitForMsg()
	s.game.Finish(winner)
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
