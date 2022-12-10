package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadLines("./input.txt")

	cycle, register, strengthSum := 0, 1, 0
	var img strings.Builder

	for _, line := range input {
		nextCycle(&cycle, &register, &strengthSum, &img)
		if line != "noop" {
			value := utils.ConvertToInt(strings.Fields(line)[1])
			nextCycle(&cycle, &register, &strengthSum, &img)
			register += value
		}
	}
	fmt.Printf("\nPart 1: %d\nPart 2:\n%s", strengthSum, img.String())
}

func nextCycle(cycle, register, strengthSum *int, image *strings.Builder) {
	draw(*cycle, *register, image)
	if *cycle++; (*cycle+20)%40 == 0 {
		*strengthSum += *cycle * *register
	}
}

func draw(cycle, register int, image *strings.Builder) {
	pos := (cycle % 40)
	if pos <= register+1 && pos >= register-1 {
		image.WriteRune('#')
	} else {
		image.WriteRune('.')
	}
	if pos == 39 {
		image.WriteRune('\n')
	}
}
