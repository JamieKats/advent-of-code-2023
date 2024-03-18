package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type srcToDest func(int) int

type mappingRange struct {
	src         int
	dest        int
	srcToDest   srcToDest
	rangeLength int
}

func main() {

	// Read input
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")

	// Parse seeds reading
	seedsRaw := split_data[0]
	seedsValuesRaw := strings.Split(seedsRaw, " ")[1:]
	var seeds []int
	for _, seed := range seedsValuesRaw {
		seedVal, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedVal)
	}

	// Contains arrays that represent the mappings for each stage, e.g.
	// mappings[0] is an array of all mappings for seed-to-soil
	// mappings[1] is an array of all mappings for soil-to-fertilizer etc
	mappings := [7][]mappingRange{}
	for i := range mappings {
		mappings[i] = []mappingRange{}
	}
	var mappingIndex int
	// Parse thing-to-thing mappings
	rawMappingsInput := split_data[2:]
	for i, line := range rawMappingsInput {
		// loop through lines and process every mapping
		if strings.Contains(line, "to") {
			// fmt.Println(line)
			// fmt.Println(i)
			// panic("23")
			// increment current mappings value
			for _, mappingRaw := range rawMappingsInput[i+1:] {
				// fmt.Println(rawMappingsInput[i+1:])
				// fmt.Println(mappingRaw)
				// fmt.Println(j)
				// panic("srdgf")

				// for _, y := range split_data[i+2:] {
				// 	fmt.Println(y)
				// }

				// fmt.Println(split_data[i+1:])
				// If line is empty break
				if mappingRaw == "" {
					mappingIndex += 1
					break
				}

				// Parse raw mapping
				mappingsSplit := strings.Split(mappingRaw, " ")
				// fmt.Println(mappingsSplit)
				start, _ := strconv.Atoi(mappingsSplit[1])
				end, _ := strconv.Atoi(mappingsSplit[0])
				srcToDest := func(x int) int { return x + (end - start) }
				rangeLength, _ := strconv.Atoi(mappingsSplit[2])

				parsedMappings := mappingRange{
					src:         start,
					dest:        end,
					srcToDest:   srcToDest,
					rangeLength: rangeLength,
				}
				// fmt.Println(mappings[mappingIndex])
				mappings[mappingIndex] = append(mappings[mappingIndex], parsedMappings)
				// fmt.Println(mappings[mappingIndex])
				// panic("wef")
			}
		}
		// fmt.Println()
		// for _, mapping := range mappings[0] {
		// 	fmt.Printf("start: %d, end: %d, func: %v\n", mapping.start, mapping.end, mapping.srcToDest)
		// }
	}

	// For each seed process seeds
	var part1Result int
	for _, seed := range seeds {
		location := processSeed(seed, &mappings)
		if location < part1Result || part1Result == 0 {
			part1Result = location
		}
	}
	fmt.Printf("Part 1: %d\n", part1Result)
}

func processSeed(seed int, mappingRangeInput *[7][]mappingRange) int {
	// pass initial seed value through all mappingRanges
	currVal := seed
	for _, mappingRange := range mappingRangeInput {
		// fmt.Println(currVal)
		for _, mapRange := range mappingRange {
			// fmt.Println(mapRange.srcToDest(currVal))
			// fmt.Println(mapRange.rangeLength)
			// panic(1)
			rangeStart := mapRange.src
			rangeEnd := mapRange.src + mapRange.rangeLength
			if currVal >= rangeStart && currVal < rangeEnd {
				currVal = mapRange.srcToDest(currVal)
				break
			}
		}
	}
	// fmt.Println(currVal)
	// panic(1)
	return currVal
}

// func sourceToDest(input int)

// func decodeCard(cardLine string) card {
// 	var decodedCard card
// 	split_data := strings.Split(string(cardLine), ":")
// 	cardNumInfo := strings.Split(split_data[0], " ")
// 	decodedCard.cardNum, _ = strconv.Atoi(cardNumInfo[len(cardNumInfo)-1])

// 	winningNumbersRawString := strings.Split(split_data[1], "|")[0]
// 	winningNumbersRawString = strings.TrimSpace(winningNumbersRawString)
// 	winningNumbersRaw := strings.Split(winningNumbersRawString, " ")
// 	winningNumbers := make([]int, len(winningNumbersRaw))

// 	// fmt.Println(winningNumbersRaw)
// 	// fmt.Println(len(winningNumbersRaw))
// 	for i, num := range winningNumbersRaw {
// 		winningNumbers[i], _ = strconv.Atoi(num)
// 	}

// 	myNumbersRawString := strings.Split(split_data[1], "|")[1]
// 	myNumbersRawString = strings.TrimSpace(myNumbersRawString)
// 	myNumbersRaw := strings.Split(myNumbersRawString, " ")
// 	myNumbers := make([]int, len(myNumbersRaw))

// 	for i, num := range myNumbersRaw {
// 		myNumbers[i], _ = strconv.Atoi(num)
// 	}
// 	// fmt.Println(myNumbersRaw)
// 	// fmt.Println(len(myNumbersRaw))
// 	decodedCard.winningNumbers = winningNumbers
// 	decodedCard.myNumbers = myNumbers

// 	// Calculate card score
// 	var score []int
// 	winningNumbers = removeSpace(winningNumbers)
// 	for _, winningNum := range winningNumbers {
// 		if containsInt(myNumbers, winningNum) {
// 			score = append(score, winningNum)
// 		}
// 	}
// 	// fmt.Printf("winning nums: %v\n", score)
// 	decodedCard.numMatchingNumbers = len(score)
// 	decodedCard.score = int(math.Pow(2, float64(len(score)-1)))
// 	return decodedCard
// }

// func containsInt(input []int, target int) bool {
// 	for _, num := range input {
// 		if num == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// func removeSpace(input []int) []int {
// 	var filteredSlice []int
// 	for _, num := range input {
// 		if num != 0 {
// 			filteredSlice = append(filteredSlice, num)
// 		}
// 	}
// 	return filteredSlice
// }

// func part2CheckCard(cards []card, cardNum int, fromCard int) int {
// 	// fmt.Println()
// 	// Get num of matches for card
// 	numMatches := cards[cardNum-1].numMatchingNumbers

// 	// base case
// 	if numMatches == 0 {
// 		// fmt.Printf("\t\t\tEnd card %d->%d\n", fromCard, cardNum)
// 		return 1
// 	}

// 	cardCopies := cards[cardNum:int(math.Min(float64(len(cards)), float64(cardNum+numMatches)))]
// 	// fmt.Printf("Card %d\n", cardNum)
// 	// fmt.Printf("Slice %d:min(%d, %d)\n", cardNum, len(cards), cardNum+numMatches)
// 	// fmt.Printf("Card copies: %v\n", cardCopies)

// 	// return 0
// 	// fmt.Printf("\t\t\tMiddle card %d->%d\n", fromCard, cardNum)
// 	copies := 1 // Current card needs to be included in the count
// 	for _, cardCopy := range cardCopies {
// 		copies += (part2CheckCard(cards, cardCopy.cardNum, cardNum)) // Current card + all current cards recursive children are counted for each copy
// 	}
// 	return copies
// }
