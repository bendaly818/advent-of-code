package main

import (
	"fmt"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

var rockPaperScissorsMapP1 = map[string]map[string]int{
	"A": {
		"X": 4,
		"Y": 8,
		"Z": 3,
	},
	"B": {
		"X": 1,
		"Y": 5,
		"Z": 9,
	},
	"C": {
		"X": 7,
		"Y": 2,
		"Z": 6,
	},
}

var rockPaperScissorsMapP2 = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 1,
		"C": 2,
	},
	"Y": {
		"A": 4,
		"B": 5,
		"C": 6,
	},
	"Z": {
		"A": 8,
		"B": 9,
		"C": 7,
	},
}

func get_p1_result(elfChoice string, myChoice string) int {
	return rockPaperScissorsMapP1[elfChoice][myChoice]
}

func get_p2_result(elfChoice string, myChoice string) int {
	return rockPaperScissorsMapP2[myChoice][elfChoice]
}

func main() {
	contents, _ := utils.ReadFile("./real.txt")

	var p1 int = 0
	var p2 int = 0

	var games = strings.Split(contents, "\n")

	for _, game := range games {
		var choices = strings.Split(game, " ")
		p1 += get_p1_result(choices[0], choices[1])
		p2 += get_p2_result(choices[0], choices[1])
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
