package main

import (
	"aoc-2022/utils"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	pairs := strings.Split(utils.ReadFile("input.txt"), "\n\n")
	var (
		p1      int
		packets []any
	)
	for i, pair := range pairs {
		pair := strings.Split(pair, "\n")

		var left, right any
		json.Unmarshal([]byte(pair[0]), &left)
		json.Unmarshal([]byte(pair[1]), &right)
		packets = append(packets, left, right)
		if compare(left, right) > 0 {
			p1 += i + 1
		}
	}

	var divider1, divider2 any
	json.Unmarshal([]byte("[[2]]"), &divider1)
	json.Unmarshal([]byte("[[6]]"), &divider2)
	packets = append(packets, divider1, divider2)

	sort.SliceStable(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) > 0
	})

	p2, found := 1, 0
	for i, p := range packets {
		if found == 2 {
			break
		}
		s, _ := json.Marshal(p)

		if string(s) == "[[2]]" || string(s) == "[[6]]" {
			p2 *= (i + 1)
			found++
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
}

func compare(left, right any) int {
	l, leftOk := left.(float64)
	r, rightOk := right.(float64)
	if leftOk && rightOk {
		return int(r) - int(l)
	}

	lArr, leftOk := left.([]any)
	rArr, rightOk := right.([]any)
	if leftOk && rightOk {
		for i := 0; i < utils.Min(len(lArr), len(rArr)); i++ {
			if r := compare(lArr[i], rArr[i]); r != 0 {
				return r
			}
		}
		return len(rArr) - len(lArr)
	}

	lArr, leftOk = left.([]any)
	rValue, rightOk := right.(float64)

	if leftOk && rightOk {
		return compare(lArr, []any{rValue})
	}

	lValue := left.(float64)
	rArr = right.([]any)

	return compare([]any{lValue}, rArr)
}
