package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type card struct {
	cardNum        int
	myNumbers      []int
	winningNumbers []int
	score          int
}

// Define generic Map function type
// type MapFunc[A any, B error] func(A) (A B)

// func Map[A any](input []A, m MapFunc[A]) []A {
// 	output := make([]A, len(input))
// 	for i, val := range input {
// 		output[i] = m(val)
// 	}
// 	return output
// }

func main() {

	// Read input
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")

	var sum int
	for _, line := range split_data {
		decodedCard := decodeCard(line)
		sum += decodedCard.score
	}
	fmt.Printf("Score: %d\n", sum)

	// var digitsFound []numInfo
	// gearsFound := make(map[string][]int)

	// for row, lineString := range split_data {
	// 	parseDigits(lineString, row, &digitsFound)
	// }

	// var sum int
	// var gearRatios int
	// for _, digit := range digitsFound {
	// 	if numIsPart(digit, split_data, gearsFound) {
	// 		sum += digit.digit
	// 	}
	// }
	// // Print result gears
	// for _, v := range gearsFound {
	// 	if len(v) == 2 {
	// 		gearRatios += v[0] * v[1]
	// 	}
	// }
	// // Final sum and gearRatios
	// fmt.Println(sum)
	// fmt.Println(gearRatios)
}

func decodeCard(cardLine string) card {
	var decodedCard card
	split_data := strings.Split(string(cardLine), ":")
	decodedCard.cardNum, _ = strconv.Atoi(strings.Split(split_data[0], " ")[1])

	winningNumbersRawString := strings.Split(split_data[1], "|")[0]
	winningNumbersRawString = strings.TrimSpace(winningNumbersRawString)
	winningNumbersRaw := strings.Split(winningNumbersRawString, " ")
	winningNumbers := make([]int, len(winningNumbersRaw))

	fmt.Println(winningNumbersRaw)
	fmt.Println(len(winningNumbersRaw))
	for i, num := range winningNumbersRaw {
		winningNumbers[i], _ = strconv.Atoi(num)
	}

	myNumbersRawString := strings.Split(split_data[1], "|")[1]
	myNumbersRawString = strings.TrimSpace(myNumbersRawString)
	myNumbersRaw := strings.Split(myNumbersRawString, " ")
	myNumbers := make([]int, len(myNumbersRaw))

	for i, num := range myNumbersRaw {
		myNumbers[i], _ = strconv.Atoi(num)
	}
	fmt.Println(myNumbersRaw)
	fmt.Println(len(myNumbersRaw))
	decodedCard.winningNumbers = winningNumbers
	decodedCard.myNumbers = myNumbers

	// Calculate card score
	var score []int
	winningNumbers = removeSpace(winningNumbers)
	for _, winningNum := range winningNumbers {
		if containsInt(myNumbers, winningNum) {
			score = append(score, winningNum)
		}
	}
	fmt.Printf("winning nums: %v\n", score)
	decodedCard.score = int(math.Pow(2, float64(len(score)-1)))
	return decodedCard
}

func containsInt(input []int, target int) bool {
	for _, num := range input {
		if num == target {
			return true
		}
	}
	return false
}

func removeSpace(input []int) []int {
	var filteredSlice []int
	for _, num := range input {
		if num != 0 {
			filteredSlice = append(filteredSlice, num)
		}
	}
	return filteredSlice
}

// func parseDigits(line string, rowNum int, digitsFound *[]numInfo) {
// 	// Check if there are digits in the line
// 	pattern := `(\d+)`

// 	re := regexp.MustCompile(pattern)

// 	indexMatches := re.FindAllStringSubmatchIndex(line, -1)
// 	stringMatches := re.FindAllStringSubmatch(line, -1)

// 	for i, matchInfo := range indexMatches {

// 		num, _ := strconv.Atoi(stringMatches[i][0])
// 		numIndex := matchInfo[0]

// 		foundNum := numInfo{
// 			digit: num,
// 			row:   rowNum,
// 			col:   numIndex,
// 		}
// 		// add found number to list
// 		*digitsFound = append(*digitsFound, foundNum)
// 	}
// }

// func numIsPart(digitFound numInfo, inputArray []string, gearsFound map[string][]int) bool {
// 	// perform check around each digit in the number
// 	for targetCol := digitFound.col; targetCol < digitFound.col+len(fmt.Sprint(digitFound.digit)); targetCol++ {

// 		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
// 			for colOffset := -1; colOffset <= 1; colOffset++ {
// 				// don't check the number if no offset
// 				if rowOffset == 0 && colOffset == 0 {
// 					continue
// 				}

// 				row := digitFound.row + rowOffset
// 				col := targetCol + colOffset

// 				// if row or col out of range of array size skip check
// 				if row < 0 || row >= len(inputArray) {
// 					continue
// 				} else if col < 0 || col >= len(inputArray[0]) {
// 					continue
// 				}

// 				// If char at row col is a digit skip
// 				if unicode.IsDigit(rune(inputArray[row][col])) {
// 					continue
// 				}

// 				// Character indicates number is a part number
// 				currChar := string(inputArray[row][col])
// 				if currChar != "." {
// 					// If we found a gear add this digit to the gears found
// 					if currChar == "*" {
// 						gearsFound[fmt.Sprintf("%d-%d", row, col)] = append(gearsFound[fmt.Sprintf("%d-%d", row, col)], digitFound.digit)
// 					}
// 					return true
// 				}
// 			}
// 		}
// 		// fmt.Println()
// 	}
// 	return false
// }
