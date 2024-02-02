package main

import (
	"fmt"
)

func main() {
	// Initialise and set a string s
	// to 'Hello World'
	s := "Hello World"

	// Write out the string to the console
	// Return the string for our later step.
	fmt.Println(messageOutput(s))
}

func messageOutput(s string) string {
	return s
}
