package main

import (
	"aoc-2022/utils"
	"fmt"
)

func main() {
	input := utils.ReadLines("./input")
	fmt.Printf("Part 1: %d\n", p1(input))
	fmt.Printf("Part 2: %d", p2(input))
}

func p1(rounds []string) (score int) {
	for _, round := range rounds {
		r := []rune(round)
		opponent, you := int(r[0]-'A')+1, int(r[2]-'X')+1
		if opponent-you == 0 {
			score += you + 3
		} else if (opponent%3)+1 == you {
			score += you + 6
		} else {
			score += you
		}
	}
	return
}

func p2(rounds []string) (score int) {
	for _, round := range rounds {
		r := []rune(round)
		switch r[2] {
		case 'X':
			score += int(r[0]-'A'+2)%3 + 1
		case 'Y':
			score += int(r[0]-'A') + 1
		case 'Z':
			score += int(r[0]-'A'+1)%3 + 1
		}
		score += int(r[2]-'X') * 3
	}
	return
}
