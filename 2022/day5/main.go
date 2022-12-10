package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

func slice[K comparable](i int, xs []K) ([]K, []K) {
	var take = i
	if take > len(xs) {
		take = len(xs)
	}
	y := xs[len(xs)-take:]
	ys := xs[0 : len(xs)-take]
	return y, ys
}

func main() {

	contents, _ := utils.ReadFile("./real.txt")

	cratesAndMoves := strings.Split(contents, "\n\n")

	var crates = cratesAndMoves[0]

	var moves = cratesAndMoves[1]
	var cratesMap = map[int][]string{}
	crateLines := strings.Split(crates, "\n")
	var crateIndexLine = crateLines[len(crateLines)-1]

	indexes := regexp.MustCompile(`\d`).FindAllStringIndex(crateIndexLine, -1)
	for i := len(crateLines) - 2; i >= 0; i-- {
		for j, crateIndex := range indexes {
			if cratesMap[j+1] == nil {
				cratesMap[j+1] = []string{}
			}

			crateLetter := string(crateLines[i][crateIndex[0]])

			if crateLetter == " " {
				continue
			}

			cratesMap[j+1] = append(cratesMap[j+1], string(crateLines[i][crateIndex[0]]))
		}
	}

	var movesArray = [][]int{}
	for _, moveLine := range strings.Split(moves, "\n") {
		moveStrings := regexp.MustCompile(`\d+`).FindAllString(moveLine, -1)
		var moveInts = []int{}
		for _, moveString := range moveStrings {
			moveInt, _ := strconv.Atoi(moveString)
			moveInts = append(moveInts, moveInt)
		}
		movesArray = append(movesArray, moveInts)
	}

	for _, moveInts := range movesArray {
		var numberToMove = moveInts[0]
		var moveFrom = moveInts[1]
		var moveTo = moveInts[2]
		take, leftovers := slice(numberToMove, cratesMap[moveFrom])
		cratesMap[moveFrom] = leftovers
		for i := len(take) - 1; i >= 0; i-- {
			cratesMap[moveTo] = append(cratesMap[moveTo], take[i])
		}

	}

	fmt.Println((len(cratesMap)))
	for i := 1; i <= len(cratesMap); i++ {
		fmt.Print(cratesMap[i][len(cratesMap[i])-1])
	}
}
