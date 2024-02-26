package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	num            int
	fewestPossible round
	isValid        bool
}

type round struct {
	red   int
	green int
	blue  int
}

var CUBE_LIMITS = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {

	// Read input
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf("ERROR READING FILE: %s\n", err)
	}

	split_data := strings.Split(string(data), "\n")

	// var possible_games_sum int
	var sum int
	var powerSum int
	for _, line := range split_data {
		// Parse line to a game struct
		decodedGame := parseLine(line)

		roundPowerSum := decodedGame.fewestPossible.blue * decodedGame.fewestPossible.red * decodedGame.fewestPossible.green
		powerSum += roundPowerSum

		if decodedGame.isValid {
			sum += decodedGame.num
		}
	}

	fmt.Printf("Result sum: %d\nResult power sum: %d\n", sum, powerSum)
}

func parseLine(line string) game {
	decodedGame := game{}
	decodedGame.isValid = true

	// split by ':'
	game_prefix := strings.Split(line, ":")
	decodedGame.num, _ = strconv.Atoi(game_prefix[0][5:])

	gameResults := strings.Split(game_prefix[1], ";")
	// fmt.Printf("game result: %s\n", gameResults)

	for _, gameResult := range gameResults {
		decodedRound := parseRoundResult(gameResult)

		validRound := checkRoundValid(decodedRound)

		if !validRound {
			decodedGame.isValid = false
		}

		// update decodedGame.fewestPossible with current round marble numbers
		if decodedRound.blue > decodedGame.fewestPossible.blue {
			decodedGame.fewestPossible.blue = decodedRound.blue
		}
		if decodedRound.red > decodedGame.fewestPossible.red {
			decodedGame.fewestPossible.red = decodedRound.red
		}
		if decodedRound.green > decodedGame.fewestPossible.green {
			decodedGame.fewestPossible.green = decodedRound.green
		}
	}

	return decodedGame
}

func parseRoundResult(roundString string) round {
	finalRound := round{}
	pattern := `(\d+)\s*(?:blue|red|green)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(roundString, -1)

	for _, matchInfo := range matches {
		match := matchInfo[0]
		matchNum := matchInfo[1]

		if strings.Contains(match, "red") {
			finalRound.red, _ = strconv.Atoi(matchNum)
		} else if strings.Contains(match, "blue") {
			finalRound.blue, _ = strconv.Atoi(matchNum)
		} else if strings.Contains(match, "green") {
			finalRound.green, _ = strconv.Atoi(matchNum)
		}
	}

	return finalRound
}

func checkRoundValid(currentRount round) bool {
	if currentRount.blue > CUBE_LIMITS["blue"] || currentRount.red > CUBE_LIMITS["red"] || currentRount.green > CUBE_LIMITS["green"] {
		return false
	}
	return true
}
