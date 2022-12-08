package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

func main() {
	contents, _ := utils.ReadFile("./real.txt")

	res := strings.Split(contents, "\n\n")

	var highestWeight int = 0
	for _, elf := range res {
		parts := strings.Split(elf, "\n")
		var totalWeight int = 0
		for _, weightString := range parts {
			weight, _ := strconv.Atoi(weightString)
			totalWeight += weight
		}
		if totalWeight > highestWeight {
			highestWeight = totalWeight
		}
	}

	fmt.Println(highestWeight)
}
