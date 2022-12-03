package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := utils.ReadLines("./input.txt")
	fmt.Printf("Part 1: %d\n", p1(input))
	fmt.Printf("Part 2: %d", p2(input))
}

func p1(rucksacks []string) (prioSum int) {
	for _, rucksack := range rucksacks {
		for _, item := range rucksack[:len(rucksack)/2] {
			if strings.ContainsRune(rucksack[len(rucksack)/2:], item) {
				if unicode.IsUpper(item) {
					prioSum += int(item-'A') + 27
				} else {
					prioSum += int(item-'a') + 1
				}
				break
			}
		}
	}
	return
}

func p2(rucksacks []string) (badgeSum int) {
	for i := 0; i < len(rucksacks)-2; i += 3 {
		for _, item := range rucksacks[i] {
			if strings.ContainsRune(rucksacks[i+1], item) && strings.ContainsRune(rucksacks[i+2], item) {
				if unicode.IsUpper(item) {
					badgeSum += int(item-'A') + 27
				} else {
					badgeSum += int(item-'a') + 1
				}
				break
			}
		}
	}
	return
}
