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

	var sum int = 0

	for _, line := range split_data {
		digitsInLine := []digitPosition{}

		// Scan line for digits
		for k, _ := range wordToDigit {
			firstLoc := strings.Index(line, k)
			lastLoc := strings.LastIndex(line, k)

			if firstLoc != -1 {
				firstDigit := digitPosition{
					digit:    wordToDigit[k],
					position: firstLoc,
				}
				digitsInLine = append(digitsInLine, firstDigit)
			}

			if lastLoc != -1 {
				lastDigit := digitPosition{
					digit:    wordToDigit[k],
					position: lastLoc,
				}
				digitsInLine = append(digitsInLine, lastDigit)
			}
		}

		// Loop through the line and save first and last numeric
		for i, char := range line {

			// If ascii digit found create digit position struct and add ti digitsInLine
			if unicode.IsNumber(char) {
				digit, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}

				d := digitPosition{
					digit:    digit,
					position: i,
				}

				digitsInLine = append(digitsInLine, d)
			}
		}
		fmt.Println(digitsInLine)

		fmt.Println(getFirstPlusLastDigitInLine(digitsInLine))
		lineResult := getFirstPlusLastDigitInLine(digitsInLine)

		sum = sum + lineResult
	}

	// Sum up all values
	// var sum int = 0
	// for _, val := range result {
	// 	convertedVal, err := strconv.Atoi(val)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	sum = sum + convertedVal
	// }
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

func getFirstPlusLastDigitInLine(digitsInLine []digitPosition) int {
	firstDigit := digitsInLine[0]
	lastDigit := digitsInLine[0]
	for _, digitInLine := range digitsInLine {
		if digitInLine.position < firstDigit.position {
			firstDigit = digitInLine
		}

		if digitInLine.position > lastDigit.position {
			lastDigit = digitInLine
		}
	}

	stringDigit := strconv.Itoa(firstDigit.digit) + strconv.Itoa(lastDigit.digit)

	result, err := strconv.Atoi(stringDigit)
	if err != nil {
		panic(err)
	}
	return result
}
