package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadLines("./input.txt")
	p1, p2 := solution(input)
	fmt.Printf("Part 1: %d\nPart 2: %d", p1, p2)
}

func solution(assignments []string) (contained, overlaps int) {
	for _, assignment := range assignments {
		start1, end1, start2, end2 := getRanges(assignment)

		if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
			contained++
		}

		if (start1 <= end2 && start2 <= end1) || (start2 <= end1 && start1 <= end2) {
			overlaps++
		}
	}
	return
}

func getRanges(a string) (int, int, int, int) {
	pairs := strings.Split(a, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")
	return utils.ConvertToInt(pair1[0]), utils.ConvertToInt(pair1[1]), utils.ConvertToInt(pair2[0]), utils.ConvertToInt(pair2[1])
}
