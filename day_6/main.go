package main

import (
	"fmt"
	"math"
)

// type srcToDest func(int) int

// type mappingRange struct {
// 	src         int
// 	dest        int
// 	srcToDest   srcToDest
// 	rangeLength int
// }

// type seedRange struct {
// 	start  int
// 	length int
// }

type lowestWinningVal func(int) int
type highestWinningVal func(int) int

type raceInfo struct {
	time              int
	dist              int
	lowestWinningVal  lowestWinningVal
	highestWinningVal highestWinningVal
}

func getLowestWinningVal(a float64, b float64, c float64) int {
	lowestVal := math.Ceil(((-float64(b) - math.Sqrt(math.Pow(b, 2)-4*a*c)) / float64((2 * a))))
	return int(lowestVal)
}

func getHighestWinningVal(a float64, b float64, c float64) int {
	highestVal := math.Floor(((-float64(b) + math.Sqrt(math.Pow(b, 2)-4*a*c)) / float64((2 * a))))
	return int(highestVal)
}

func main() {

	x := getLowestWinningVal(1, -7, 9)
	y := getHighestWinningVal(1, -7, 9)
	for i := x; i <= y; i++ {
		fmt.Println(i)
	}
	panic(1)

	// Read input
	// data, err := os.ReadFile("./input.txt")

	// if err != nil {
	// 	fmt.Printf("ERROR READING FILE: %s\n", err)
	// }

	// split_data := strings.Split(string(data), "\n")

	// // Parse seeds reading
	// seedsRaw := split_data[0]
	// seedsValuesRaw := strings.Split(seedsRaw, " ")[1:]
	// var seeds []int

}
