package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

func main() {
	contents, _ := utils.ReadFile("./real.txt")

	res := strings.Split(contents, "\n\n")

	weights := []int{}
	for _, elf := range res {
		parts := strings.Split(elf, "\n")
		var totalWeight int = 0
		for _, weightString := range parts {
			weight, _ := strconv.Atoi(weightString)
			totalWeight += weight
		}
		weights = append(weights, totalWeight)
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] > weights[j]
	})

	fmt.Println(weights[0] + weights[1] + weights[2])
}
