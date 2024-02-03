package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type digitPosition struct {
	digit    int
	position int
}

func main() {
	data, err := os.ReadFile("./input.txt")

	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")

	var result []string

	for _, line := range split_data {
		// Scan line for digits
		for k, v := range wordToDigit {
			loc := strings.Index(line, k)
		}

		// Loop through the line and save first and last value
		var digitsInLine []string
		for _, char := range line {
			// fmt.Println(string(char))
			if unicode.IsNumber(char) {
				digitsInLine = append(digitsInLine, string(char))
			}
		}
		// fmt.Printf("%s%s\n", digitsInLine[0], digitsInLine[len(digitsInLine)-1])
		var lineResult = digitsInLine[0] + digitsInLine[len(digitsInLine)-1]
		result = append(result, lineResult)
	}

	// Sum up all values
	var sum int = 0
	for _, val := range result {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		sum = sum + convertedVal
	}
	fmt.Printf("Sum: %d\n", sum)

	// Write output to file
	f, err := os.Create("out.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, value := range result {
		fmt.Fprintln(f, value)
	}
}
