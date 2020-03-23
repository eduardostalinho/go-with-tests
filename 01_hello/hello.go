package main

import "fmt"

const prefix = "Hello"

// Hello receives <name> and returns "Hello <name>!"
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s!", prefix, name)
}

func main() {
	fmt.Println(Hello(""))
}
