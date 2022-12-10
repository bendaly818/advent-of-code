package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

func main() {
	contents, _ := utils.ReadFile("./real.txt")

	var duplicatesPartOne = 0
	var duplicatesPartTwo = 0

	for _, line := range strings.Split(contents, "\n") {
		leftAndRight := strings.Split(line, ",")
		left := leftAndRight[0]
		leftBottom, _ := strconv.Atoi(strings.Split(left, "-")[0])
		leftTop, _ := strconv.Atoi(strings.Split(left, "-")[1])
		right := leftAndRight[1]
		rightBottom, _ := strconv.Atoi(strings.Split(right, "-")[0])
		rightTop, _ := strconv.Atoi(strings.Split(right, "-")[1])

		if (leftBottom >= rightBottom && leftTop <= rightTop) || (leftBottom <= rightBottom && leftTop >= rightTop) {
			duplicatesPartOne += 1
		}

		if (leftTop >= rightBottom && rightTop >= leftTop) || (rightTop >= leftBottom && leftTop >= rightTop) {
			duplicatesPartTwo += 1
		}
	}

	fmt.Println(duplicatesPartOne)
	fmt.Println(duplicatesPartTwo)
}
