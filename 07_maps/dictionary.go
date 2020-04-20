package main

import "errors"

var ErrUnexistingWord = errors.New("word not found in dictionary")

type Dictionary map[string]string

func (d Dictionary) Add(w, definition string) {
	d[w] = definition
}

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrUnexistingWord
	}
	return definition, nil
}
