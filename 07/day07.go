package main

import (
	"aoc-2022/utils"
	"fmt"
	"math"
	"path"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadLinesSplit("./input.txt", '\n')
	fs := make(map[string]int)
	cwd := "/"

	for _, line := range input[1:] {
		args := strings.Fields(strings.TrimSpace(line))
		if args[0] == "$" {
			if args[1] == "cd" {
				cwd = path.Join(cwd, args[2])
			}
			continue
		}
		if fileSize, err := strconv.Atoi(args[0]); err == nil {
			dir := cwd
			for {
				fs[dir] += fileSize
				if dir == "/" {
					break
				}
				dir = path.Dir(dir)
			}
		}
	}
	part1, part2 := 0, math.MaxInt
	space := 70000000 - fs["/"] - 30000000
	for _, size := range fs {
		if size <= 100000 {
			part1 += size
		}
		if size+space > 0 && size < part2 {
			part2 = size
		}
	}
	fmt.Printf("Part 1: %d\nPart 2: %d", part1, part2)
}
