package main

import (
	"fmt"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

var lowercaseCode = int("a"[0])

func get_char_value(char rune) int {
	charAsInt := int(char)
	if charAsInt >= lowercaseCode {
		return charAsInt - 96
	}

	return charAsInt - 38
}

func p1(bags []string) int {
	total := 0

	for _, bag := range bags {
		numItems := len(bag)
		half := numItems / 2
		first := bag[0:half]
		second := bag[half:numItems]

		for _, char := range first {
			if strings.ContainsRune(second, char) {
				total += get_char_value(char)
				break
			}
		}
	}

	return total
}

func p2(bags []string) int {
	total := 0

	for i, bag := range bags {
		if i%3 != 0 {
			continue
		}
		for _, char := range bag {
			if strings.ContainsRune(bags[i+1], char) && strings.ContainsRune(bags[i+2], char) {
				total += get_char_value(char)
				break
			}
		}
	}

	return total
}

func main() {

	contents, _ := utils.ReadFile("./real.txt")

	bags := strings.Split(contents, "\n")

	fmt.Println(p1(bags))
	fmt.Println(p2(bags))
}
