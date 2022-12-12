package main

import (
	"aoc-2022/utils"
	"fmt"
	"image"
	"math"
)

func main() {
	input := utils.ReadLines("input.txt")
	grid, start, end := createGrid(input)
	p1, p2 := shortestPaths(grid, start, end)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
}

func createGrid(input []string) (grid [][]rune, start, end image.Point) {
	grid = make([][]rune, len(input))
	for y, r := range input {
		row := make([]rune, len(r))
		for x, c := range r {
			if c == 'S' {
				c = 'a'
				start = image.Point{x, y}
			} else if c == 'E' {
				c = 'z'
				end = image.Point{x, y}
			}
			row[x] = c
		}
		grid[y] = row
	}
	return
}

func adjacent(grid [][]rune, p image.Point) []image.Point {
	var result []image.Point
	dirs := []image.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, dir := range dirs {
		nx, ny := p.X+dir.X, p.Y+dir.Y
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[0]) {
			if grid[p.Y][p.X]-grid[ny][nx] <= 1 {
				result = append(result, image.Point{nx, ny})
			}
		}
	}
	return result

}

type node struct {
	point image.Point
	steps int
}

func shortestPaths(grid [][]rune, start, end image.Point) (int, int) {
	q := []node{{end, 0}}
	e := map[image.Point]struct{}{}
	p1, p2 := 0, math.MaxInt

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if grid[current.point.Y][current.point.X] == 'a' && current.steps < p2 {
			p2 = current.steps
		}
		if current.point.Eq(start) {
			p1 = current.steps
			break
		}
		for _, a := range adjacent(grid, current.point) {
			if _, ok := e[a]; !ok {
				e[a] = struct{}{}
				q = append(q, node{a, current.steps + 1})
			}
		}
	}

	return p1, p2
}
