package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func LeagueFromReader(r io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		_ = fmt.Errorf("error parsing json database. %v", err)
	}
	return league, err
}
