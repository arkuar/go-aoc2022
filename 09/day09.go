package main

import (
	"aoc-2022/utils"
	"fmt"
	"image"
	"strings"
)

const ROPESIZE = 10

var dirs = map[string]image.Point{"R": {1, 0}, "L": {-1, 0}, "U": {0, -1}, "D": {0, 1}}

func main() {
	input := utils.ReadLines("./input.txt")

	v1 := make(map[image.Point]struct{})
	v2 := make(map[image.Point]struct{})

	rope := make([]image.Point, ROPESIZE)
	head := &rope[0]

	for _, line := range input {
		motion := strings.Fields(line)
		dir, steps := motion[0], utils.ConvertToInt(motion[1])

		for i := 0; i < steps; i++ {
			*head = head.Add(dirs[dir])
			for p, knot := range rope[1:] {
				rope[p+1] = moveKnot(knot, rope[p])
			}
			v1[rope[1]] = struct{}{}
			v2[rope[len(rope)-1]] = struct{}{}
		}
	}
	fmt.Println(len(v1))
	fmt.Println(len(v2))
}

func moveKnot(knot, parent image.Point) image.Point {
	if dist := parent.Sub(knot); utils.Abs(dist.X) > 1 || utils.Abs(dist.Y) > 1 {
		return knot.Add(image.Point{utils.Sgn(dist.X), utils.Sgn(dist.Y)})
	}
	return knot
}
