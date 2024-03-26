package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// type lowestWinningVal func(int) int
// type highestWinningVal func(int) int

// type raceInfo struct {
// 	time              int
// 	dist              int
// 	lowestWinningVal  lowestWinningVal
// 	highestWinningVal highestWinningVal
// }

type raceInfo struct {
	raceNum        int
	firstValidTime int
	lastValidTime  int
}

func getLowestWinningVal(a float64, b float64, c float64) int {
	lowestFloat := ((-float64(b) - math.Sqrt(math.Pow(b, 2)-4*a*c)) / float64((2 * a)))
	lowestVal := math.Floor(lowestFloat + 1)
	return int(lowestVal)
}

func getHighestWinningVal(a float64, b float64, c float64) int {
	highestFloat := ((-float64(b) + math.Sqrt(math.Pow(b, 2)-4*a*c)) / float64((2 * a)))
	highestVal := math.Ceil(highestFloat - 1)
	return int(highestVal)
}

// Trick to this one is that you can create a quadratic formula from the input
// then simply solve for the x intercepts
//
// h^2 - lh + D < 0
// Where:
//
//	h = number of milisecods held (unknown)
//	l = length of race in miliseconds (given)
//	D = record distance (given)
func main() {

	// Read input
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")
	timeLine := split_data[0]
	distanceLine := split_data[1]

	rawTime := strings.Split(timeLine, " ")
	rawDistance := strings.Split(distanceLine, " ")

	// Remove first word from raw data
	rawTime = rawTime[1:]
	rawDistance = rawDistance[1:]
	// fmt.Println(rawTime)
	// fmt.Println(rawDistance)

	// clean spaces from input
	rawTime = cleanSpaces(rawTime)
	rawDistance = cleanSpaces(rawDistance)

	// calculate range of valid answers
	raceAnswers := []raceInfo{}
	for i := range rawTime {
		time, _ := strconv.Atoi(rawTime[i])
		dist, _ := strconv.Atoi(rawDistance[i])
		rangeStart := getLowestWinningVal(1, -float64(time), float64(dist))
		rangeEnd := getHighestWinningVal(1, -float64(time), float64(dist))
		// fmt.Println(rangeStart)

		currentRace := raceInfo{
			raceNum:        i,
			firstValidTime: rangeStart,
			lastValidTime:  rangeEnd,
		}

		raceAnswers = append(raceAnswers, currentRace)
	}

	part1Answer := 1
	for _, race := range raceAnswers {
		part1Answer *= race.lastValidTime - race.firstValidTime + 1
	}
	fmt.Printf("Part 1: %d\n", part1Answer)

	// // Parse seeds reading
	// seedsRaw := split_data[0]
	// seedsValuesRaw := strings.Split(seedsRaw, " ")[1:]
	// var seeds []int

}

func cleanSpaces(input []string) []string {
	var cleanedSlice []string
	for _, value := range input {
		if value != "" {
			cleanedSlice = append(cleanedSlice, value)
		}
	}
	return cleanedSlice
}
