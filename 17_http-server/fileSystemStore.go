package poker

import (
	"encoding/json"
	"fmt"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

type FileSystemStore struct {
	encoder *json.Encoder
	league  League
}

func initDBfile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file %s from file system, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initDBfile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing db file %s, %v", file.Name(), err)
	}

	league, err := LeagueFromReader(file)

	if err != nil {
		return nil, fmt.Errorf("Problem loading league from file %s, %v", file.Name(), err)
	}

	encoder := json.NewEncoder(&tape{file})
	return &FileSystemStore{encoder, league}, nil
}

func (s *FileSystemStore) GetLeague() League {
	return s.league
}

func (s *FileSystemStore) RecordWin(name string) {
	player := s.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}
	s.encoder.Encode(s.league)

}

func (s *FileSystemStore) GetPlayerScore(name string) int {
	player := s.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}
