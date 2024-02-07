package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

// Hello returns a static string of Hello, world to the caller.
func Hello(name string) string {
	return "Hello, " + name
}
