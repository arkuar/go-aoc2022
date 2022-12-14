package main

import (
	"aoc-2022/utils"
	"fmt"
	"image"
	"math"
	"strings"
)

func main() {
	paths := utils.ReadLines("input.txt")
	grid := map[image.Point]struct{}{}

	bottom := math.MinInt
	for _, path := range paths {
		coords := strings.Split(path, " -> ")
		rocks := []image.Point{}
		for _, c := range coords {
			xy := utils.ParseInt(strings.Split(c, ","), 10)
			rocks = append(rocks, image.Pt(xy[0], xy[1]))
		}
		for prev, rock := range rocks[1:] {
			grid[rock] = struct{}{}
			if rock.Y > bottom {
				bottom = rock.Y
			}
			d := rocks[prev].Sub(rock)
			dir := image.Pt(utils.Sgn(d.X), utils.Sgn(d.Y))
			for i := 0; i < utils.Max(utils.Abs(d.X), utils.Abs(d.Y)); i++ {
				rock = rock.Add(dir)
				grid[rock] = struct{}{}
				if rock.Y > bottom {
					bottom = rock.Y
				}
			}
		}
	}
	p1, p2 := dropSand(grid, bottom)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", p1, p2)

}

func dropSand(grid map[image.Point]struct{}, bottom int) (p1, p2 int) {
	dirs := []image.Point{image.Pt(0, 1), image.Pt(-1, 1), image.Pt(1, 1)}
	pos := image.Pt(500, 0)
	for {
		if _, ok := grid[image.Pt(500, 0)]; ok {
			return
		}
		for i, dir := range dirs {
			if _, ok := grid[pos.Add(dir)]; !ok && pos.Add(dir).Y < bottom+2 {
				pos = pos.Add(dir)
				break
			}
			if i == len(dirs)-1 {
				if pos.Y >= bottom && p1 == 0 {
					p1 = p2
				}
				grid[pos] = struct{}{}
				pos = image.Pt(500, 0)
				p2++
			}
		}
	}
}
