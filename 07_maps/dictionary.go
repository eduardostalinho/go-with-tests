package main

type Dictionary map[string]string

func Search(d map[string]string, s string) string {
	return d[s]
}
