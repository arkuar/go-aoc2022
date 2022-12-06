package main

import (
	"aoc-2022/utils"
	"fmt"
	"math"
)

func main() {
	input := utils.ReadFile("./input.txt")

	startOfPacket := findStartOf(input, 4)
	startOfMsg := findStartOf(input, 14)

	fmt.Printf("Part 1: %d\nPart 2: %d", startOfPacket, startOfMsg)
}

func findStartOf(sequence string, distinctCount int) int {
	for i := 0; i < len(sequence)-distinctCount-1; i++ {
		occurences := make(map[byte]struct{}, 0)
		occurences[sequence[i]] = struct{}{}

		for j := 0; j < distinctCount-1; j++ {
			occurences[sequence[i+j+1]] = struct{}{}
		}
		if len(occurences) == distinctCount {
			return i + distinctCount
		}
	}
	return math.MinInt
}
