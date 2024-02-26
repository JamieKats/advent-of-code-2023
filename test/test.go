package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// Input string
	str := "3 green, 15 blue, 14 red"

	// Define the regular expression pattern
	pattern := `(\d+)\s*(?:blue|red|green)`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(str, -1)

	// Map to store the results
	colorNumbers := make(map[string]int)

	// Iterate over matches
	for _, match := range matches {
		// Extract color and number
		color := match[1]
		number, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}

		// Store in the map
		colorNumbers[color] = number
	}

	// Print the results
	fmt.Println("Numbers associated with colors:")
	fmt.Println("Blue:", colorNumbers["blue"])
	fmt.Println("Red:", colorNumbers["red"])
	fmt.Println("Green:", colorNumbers["green"])
}
