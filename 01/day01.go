package main

import (
	"aoc-2022/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := strings.Split(utils.ReadFile("./input"), "\n\n")
	calories := countCalories(input)
	fmt.Printf("Part 1: %d\n", calories[len(calories)-1])
	fmt.Printf("Part 2: %d", utils.SumSlice(calories[len(calories)-3:]))
}

func countCalories(input []string) []int {
	result := make([]int, len(input))
	for i, row := range input {
		for _, calories := range strings.Split(row, "\n") {
			if calories != "" {
				result[i] += utils.ConvertToInt(calories)
			}
		}
	}
	sort.Ints(result)
	return result
}
