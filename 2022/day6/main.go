package main

import (
	"fmt"
	"strings"

	"github.com/bendaly818/advent-of-code/2022/utils"
)

func main() {
	contents, _ := utils.ReadFile("./real.txt")

	letters := strings.Split(contents, "")

	var marker = ""

	for i, char := range letters {
		if len(marker) == 14 {
			fmt.Println(i)
			break
		}

		if strings.Contains(marker, char) {
			fmt.Println(marker, char)
			marker = marker[strings.Index(marker, char)+1:] + char
			continue
		}

		marker += char
	}
	fmt.Print(marker)
}
