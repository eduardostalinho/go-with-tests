package poker

import (
	bolt "go.etcd.io/bbolt"
	"strconv"
)

type BoltPlayerStore struct {
	db         *bolt.DB
	bucketName string
}

func NewBoltPlayerStore(path, bucket string) *BoltPlayerStore {
	db, _ := bolt.Open(path, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	return &BoltPlayerStore{db, bucket}
}

func (s *BoltPlayerStore) RecordWin(name string) {
	s.db.Update(func(tx *bolt.Tx) error {
		playerName := []byte(name)
		b := tx.Bucket([]byte(s.bucketName))
		value := b.Get(playerName)
		intValue, _ := strconv.Atoi(string(value))
		intValue++
		err := b.Put(playerName, []byte(strconv.Itoa(intValue)))
		return err
	})
}

func (s *BoltPlayerStore) GetPlayerScore(name string) int {
	scoreChannel := make(chan int, 1)
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.bucketName))
		v := b.Get([]byte(name))
		score, _ := strconv.Atoi(string(v))
		scoreChannel <- score
		return nil
	})
	score := <-scoreChannel
	return score
}

func (s *BoltPlayerStore) GetLeague() League {
	league := League{}
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			score, _ := strconv.Atoi(string(v))
			player := Player{string(k), score}
			league = append(league, player)
		}
		return nil
	})
	return league
}

func (s *BoltPlayerStore) Close() {
	s.db.Close()
}
