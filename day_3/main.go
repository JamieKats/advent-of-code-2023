package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type numInfo struct {
	digit int
	row   int
	col   int
}

func main() {

	// Read input
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")

	var digitsFound []numInfo
	gearsFound := make(map[string][]int)

	for row, lineString := range split_data {
		parseDigits(lineString, row, &digitsFound)
	}

	var sum int
	var gearRatios int
	for _, digit := range digitsFound {
		if numIsPart(digit, split_data, gearsFound) {
			sum += digit.digit
		}
	}
	// Print result gears
	for _, v := range gearsFound {
		if len(v) == 2 {
			gearRatios += v[0] * v[1]
		}
	}
	// Final sum and gearRatios
	fmt.Println(sum)
	fmt.Println(gearRatios)
}

func parseDigits(line string, rowNum int, digitsFound *[]numInfo) {
	// Check if there are digits in the line
	pattern := `(\d+)`

	re := regexp.MustCompile(pattern)

	indexMatches := re.FindAllStringSubmatchIndex(line, -1)
	stringMatches := re.FindAllStringSubmatch(line, -1)

	for i, matchInfo := range indexMatches {

		num, _ := strconv.Atoi(stringMatches[i][0])
		numIndex := matchInfo[0]

		foundNum := numInfo{
			digit: num,
			row:   rowNum,
			col:   numIndex,
		}
		// add found number to list
		*digitsFound = append(*digitsFound, foundNum)
	}
}

func numIsPart(digitFound numInfo, inputArray []string, gearsFound map[string][]int) bool {
	// perform check around each digit in the number
	for targetCol := digitFound.col; targetCol < digitFound.col+len(fmt.Sprint(digitFound.digit)); targetCol++ {

		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
			for colOffset := -1; colOffset <= 1; colOffset++ {
				// don't check the number if no offset
				if rowOffset == 0 && colOffset == 0 {
					continue
				}

				row := digitFound.row + rowOffset
				col := targetCol + colOffset

				// if row or col out of range of array size skip check
				if row < 0 || row >= len(inputArray) {
					continue
				} else if col < 0 || col >= len(inputArray[0]) {
					continue
				}

				// If char at row col is a digit skip
				if unicode.IsDigit(rune(inputArray[row][col])) {
					continue
				}

				// Character indicates number is a part number
				currChar := string(inputArray[row][col])
				if currChar != "." {
					// If we found a gear add this digit to the gears found
					if currChar == "*" {
						gearsFound[fmt.Sprintf("%d-%d", row, col)] = append(gearsFound[fmt.Sprintf("%d-%d", row, col)], digitFound.digit)
					}
					return true
				}
			}
		}
		// fmt.Println()
	}
	return false
}
