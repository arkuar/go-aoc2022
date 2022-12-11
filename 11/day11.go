package main

import (
	"aoc-2022/utils"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Monkey struct {
	items                                             []int
	op                                                string
	opVal, testVal, testTrue, testFalse, inspectCount int
}

func (m *Monkey) Test(x int) int {
	if x%m.testVal == 0 {
		return m.testTrue
	}
	return m.testFalse
}

func (m *Monkey) Inspect() (int, bool) {
	if len(m.items) == 0 {
		return -1, false
	}
	x := m.items[0]
	m.items = m.items[1:]
	switch m.op {
	case "+":
		x += m.opVal
	case "*":
		x *= m.opVal
	default:
		x *= x
	}
	m.inspectCount++
	return x, true
}

func main() {
	input := strings.Split(utils.ReadFile("./input.txt"), "\n\n")

	monkeys, lcm := createMonkeys(input)

	fmt.Printf("Part 1: %d\nPart 2: %d\n",
		simulate(monkeys, 20, func(i int) int { return i / 3 }),
		simulate(monkeys, 10000, func(i int) int {
			return i % lcm
		}))

}

func simulate(original []Monkey, rounds int, worryFunc func(int) int) int {
	monkeys := make([]Monkey, len(original))
	copy(monkeys, original)

	for r := 0; r < rounds; r++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for level, ok := monkey.Inspect(); ok; level, ok = monkey.Inspect() {
				level = worryFunc(level)
				monkeys[monkey.Test(level)].items = append(monkeys[monkey.Test(level)].items, level)
			}
		}
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})
	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

const noteTemplate = `Monkey %d:
		Starting items: %s
		Operation: new = old %s %d
		Test: divisible by %d
		  If true: throw to monkey %d
		  If false: throw to monkey %d`

func createMonkeys(input []string) ([]Monkey, int) {
	var monkeys []Monkey
	var lcm int
	for _, p := range input {
		monkey := Monkey{}

		var (
			items, op                       string
			id, operand, test, testT, testF int
		)

		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "** 0").Replace(p), noteTemplate, &id, &items, &op, &operand, &test, &testT, &testF)

		json.Unmarshal([]byte("["+items+"]"), &monkey.items)
		monkey.op, monkey.opVal, monkey.testVal, monkey.testTrue, monkey.testFalse = op, operand, test, testT, testF
		monkeys = append(monkeys, monkey)
		if lcm == 0 {
			lcm = test
		} else {
			lcm = utils.LCM(lcm, test)
		}

	}
	return monkeys, lcm
}
