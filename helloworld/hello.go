package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

const englishHelloPrefix = "Hello, "

// Hello returns a static string of Hello, world to the caller.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
