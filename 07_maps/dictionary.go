package main

type Dictionary map[string]string

func (d Dictionary) Search(s string) string {
	return d[s]
}
