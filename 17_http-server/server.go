package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := GetPlayerScore(player)
	fmt.Fprintf(w, score)
}

func GetPlayerScore(player string) string {
	var score string
	switch player {
	case "Adam":
		score = "20"
	case "Alice":
		score = "10"
	}
	return score

}
