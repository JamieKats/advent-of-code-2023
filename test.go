package main

import "testing"

func TestOutput(t *testing.T) {
	expected := "Hello World!"
	if actual := messageOutput("Hello World!"); actual != expected {
		t.Errorf("The main function provided an actual of %q but %q was expected.", actual, expected)
	}
}
