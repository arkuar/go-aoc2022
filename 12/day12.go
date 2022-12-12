package main

import (
	"aoc-2022/utils"
	"fmt"
	"image"
	"math"
)

func main() {
	input := utils.ReadLines("input.txt")
	grid, start, end, start2 := createGrid(input)
	p1 := shortestPath(grid, start, end)

	p2 := math.MaxInt
	for _, s := range start2 {
		res := shortestPath(grid, s, end)
		if res != -1 && res < p2 {
			p2 = res
		}
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)
}

func createGrid(input []string) (grid [][]rune, start, end image.Point, start2 []image.Point) {
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
			} else if c == 'a' {
				start2 = append(start2, image.Point{x, y})
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
			if grid[ny][nx]-grid[p.Y][p.X] <= 1 {
				result = append(result, image.Point{nx, ny})
			}
		}
	}
	return result

}

func shortestPath(grid [][]rune, start, end image.Point) int {
	q := []image.Point{start}
	e := map[image.Point]image.Point{}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current.Eq(end) {
			break
		}
		for _, a := range adjacent(grid, current) {
			if _, ok := e[a]; !ok {
				e[a] = current
				q = append(q, a)
			}
		}
	}

	if _, ok := e[end]; !ok {
		return -1
	}

	var p1 int
	for !end.Eq(start) {
		p1++
		end = e[end]
	}
	return p1
}
