package main

import (
	"aoc-2022/utils"
	"fmt"
	"image"
	"math"
)

const (
	target        = 2000000
	space         = 4000000
	inputTemplate = `Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d`
)

type Point struct {
	sensor, beacon image.Point
}

type Quadrant struct {
	min, max image.Point
}

func main() {
	input := utils.ReadLines("input.txt")
	var points []Point
	for _, line := range input {
		var sx, sy, bx, by int
		fmt.Sscanf(line, inputTemplate, &sx, &sy, &bx, &by)
		points = append(points, Point{image.Pt(sx, sy), image.Pt(bx, by)})
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(points), part2(points))
}

func part1(points []Point) (count int) {
	minX, maxX := math.MaxInt, math.MinInt

	for _, p := range points {
		minX = utils.Min(p.sensor.X-utils.Manhattan(p.sensor, p.beacon), minX)
		maxX = utils.Max(p.sensor.X+utils.Manhattan(p.sensor, p.beacon), maxX)
	}

	for x := minX; x < maxX; x++ {
		pos := image.Pt(x, target)
		for _, p := range points {
			d := utils.Manhattan(p.sensor, pos)
			if d <= utils.Manhattan(p.sensor, p.beacon) && !p.beacon.Eq(image.Pt(x, target)) {
				count++
				break
			}
		}
	}
	return
}

func canDiscard(p Point, quadrant Quadrant) bool {
	maxDistance := math.MinInt
	for _, q := range []image.Point{quadrant.min, quadrant.max, image.Pt(quadrant.min.X, quadrant.max.Y), image.Pt(quadrant.max.X, quadrant.min.Y)} {
		maxDistance = utils.Max(maxDistance, utils.Manhattan(p.sensor, q))
	}
	return maxDistance <= utils.Manhattan(p.sensor, p.beacon)
}

func splitQuadrant(q Quadrant) []Quadrant {
	min, max := q.min, q.max

	mid := image.Point{
		(min.X + max.X) / 2,
		(min.Y + max.Y) / 2,
	}
	quadrants := []Quadrant{
		{min, mid},
		{mid.Add(image.Pt(1, 1)), max},
		{image.Pt(mid.X+1, min.Y), image.Pt(max.X, mid.Y)},
		{image.Pt(min.X, mid.Y+1), image.Pt(mid.X, max.Y)},
	}

	var result []Quadrant
	for _, quad := range quadrants {
		if quad.min.X <= quad.max.X && quad.min.Y <= quad.max.Y {
			result = append(result, quad)
		}
	}
	return result
}

func part2(points []Point) int {
	quadrants := utils.Stack[Quadrant]{}
	quadrants.Push(Quadrant{
		image.Pt(0, 0),
		image.Pt(space, space),
	})

outer:
	for quadrant, ok := quadrants.Pop(); ok; quadrant, ok = quadrants.Pop() {
		for _, p := range points {
			if canDiscard(p, quadrant) {
				continue outer
			}
		}
		if quadrant.min.Eq(quadrant.max) {
			return quadrant.min.X*4000000 + quadrant.min.Y
		}
		quadrants.PushN(splitQuadrant(quadrant))
	}
	return -1
}
