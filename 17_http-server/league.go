package main
import (
	"encoding/json"
)
type League []Player


func LeagueFromReader(r io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		_ = fmt.Errorf("error parsing json database. %v", err)
	}
	return league, err
}
