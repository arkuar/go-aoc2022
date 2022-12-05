package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
)

func main() {
	input := strings.Split(utils.ReadFile("./input.txt"), "\n\n")
	data := strings.Split(input[0], "\n")
	instructions := input[1]

	num := (len(data[0]) + 1) / 4
	stacks := make([]utils.Stack[rune], num)
	stacks9001 := make([]utils.Stack[rune], num)

	for j := len(data) - 2; j >= 0; j-- {
		row := data[j]
		for i := 0; i < num; i++ {
			crate := row[i*4+1]
			if crate != ' ' {
				stacks[i].Push(rune(crate))
				stacks9001[i].Push(rune(crate))
			}
		}
	}

	for _, row := range strings.Split(strings.TrimSpace(instructions), "\n") {
		var amount, from, to int
		fmt.Sscanf(row, "move %d from %d to %d", &amount, &from, &to)

		// Part1
		for i := 0; i < amount; i++ {
			crate, _ := stacks[from-1].Pop()
			stacks[to-1].Push(crate)
		}

		// Part 2
		crates, _ := stacks9001[from-1].PopN(amount)
		stacks9001[to-1].PushN(crates)
	}

	var p1, p2 string
	for i := 0; i < num; i++ {
		s1, _ := stacks[i].Peek()
		s2, _ := stacks9001[i].Peek()
		p1 += string(s1)
		p2 += string(s2)
	}
	fmt.Printf("Part 1: %s\nPart 2: %s", p1, p2)

}
