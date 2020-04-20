package main

import "fmt"

const englishPrefix = "Hello"
const spanishPrefix = "Hola"
const germanPrefix = "Hallo"

// Greet receives <name> and <language> and returns a
// greet for the <name> in the given language
func Greet(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefix
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "german":
		prefix = germanPrefix
	}
	return fmt.Sprintf("%s %s!", prefix, name)
}

func main() {
	fmt.Println(Greet("", "english"))
}
