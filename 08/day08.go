package main

import (
	"aoc-2022/utils"
	"fmt"
	"math"
)

func main() {
	input := utils.ReadLines("./input.txt")

	grid := make([][]int, len(input))

	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, height := range line {
			grid[i][j] = utils.ConvertToInt(string(height))
		}
	}

	visible := (len(grid) + len(grid[0]) - 2) * 2
	highestScore := math.MinInt
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			ok, score := isVisible(y, x, grid)
			if ok {
				visible++
			}
			if score > highestScore {
				highestScore = score
			}
		}
	}
	fmt.Println(visible)
	fmt.Println(highestScore)
}

func isVisible(y, x int, grid [][]int) (bool, int) {
	visibleRight, visibleLeft, visibleUp, visibleDown := true, true, true, true
	scores := make([]int, 4)

	for i := len(grid[y][:x]) - 1; i >= 0; i-- {
		scores[0]++
		if grid[y][:x][i] >= grid[y][x] {
			visibleLeft = false
			break
		}
	}

	for _, right := range grid[y][x+1:] {
		scores[1]++
		if right >= grid[y][x] {
			visibleRight = false
			break
		}
	}

	for i := len(grid[:y]) - 1; i >= 0; i-- {
		scores[2]++
		if grid[:y][i][x] >= grid[y][x] {
			visibleUp = false
			break
		}
	}

	for _, down := range grid[y+1:] {
		scores[3]++
		if down[x] >= grid[y][x] {
			visibleDown = false
			break
		}
	}

	return visibleDown || visibleLeft || visibleRight || visibleUp, utils.MulSlice(scores)
}
