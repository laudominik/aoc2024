package main

import (
	"fmt"
	"os"
	"strings"
)

type Vec2i struct {
	y int
	x int
}

type PositionDirection struct {
	pos Vec2i
	dir Vec2i
}

type stack []Vec2i

func (s stack) Push(v Vec2i) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, Vec2i) {
	if len(s) == 0 {
		panic(0)
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

func parse() (gridWithBorders []string) {
	data, _ := os.ReadFile("data/day12.txt")
	grid := strings.Split(strings.TrimSpace(string(data)), "\n")

	upperBorder := ".."
	for _, _ = range grid[0] {
		upperBorder += "."
	}
	gridWithBorders = append(gridWithBorders, upperBorder)
	for _, row := range grid {
		s := "."
		for _, c := range row {
			s += string(c)
		}
		s += "."
		gridWithBorders = append(gridWithBorders, s)
	}
	gridWithBorders = append(gridWithBorders, upperBorder)
	return gridWithBorders
}

func solve1(grid []string) {
	score1 := 0
	score2 := 0
	visited := make(map[Vec2i]bool)
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if visited[Vec2i{y, x}] {
				continue
			}
			tag := grid[y][x]
			var st stack
			var dirSt stack
			st = st.Push(Vec2i{y, x})
			dirSt = dirSt.Push(Vec2i{y, x})
			area := 0
			visitedPerimeter := make(map[PositionDirection]bool)
			inPerimeter := make(map[PositionDirection]bool)

			visitedLocal := make(map[Vec2i]bool)
			for len(st) != 0 {
				var pos Vec2i
				var dir Vec2i
				st, pos = st.Pop()
				dirSt, dir = dirSt.Pop()
				if visitedLocal[pos] {
					continue
				}
				if grid[pos.y][pos.x] != tag {
					visitedPerimeter[PositionDirection{pos, dir}] = false
					inPerimeter[PositionDirection{pos, dir}] = true
					continue
				}
				visited[pos] = true
				visitedLocal[pos] = true
				st = st.Push(Vec2i{pos.y - 1, pos.x})
				dirSt = dirSt.Push(Vec2i{-1, 0})
				st = st.Push(Vec2i{pos.y + 1, pos.x})
				dirSt = dirSt.Push(Vec2i{1, 0})
				st = st.Push(Vec2i{pos.y, pos.x - 1})
				dirSt = dirSt.Push(Vec2i{0, -1})
				st = st.Push(Vec2i{pos.y, pos.x + 1})
				dirSt = dirSt.Push(Vec2i{0, 1})
				area++
			}

			sides := 0

			for k := range inPerimeter {
				if visitedPerimeter[k] {
					continue
				}
				sides++

				pos := k.pos
				dir := k.dir

				for inPerimeter[PositionDirection{pos, dir}] {
					visitedPerimeter[PositionDirection{pos, dir}] = true
					pos.y -= dir.x
					pos.x += dir.y
				}

				pos = k.pos
				for inPerimeter[PositionDirection{pos, dir}] {
					visitedPerimeter[PositionDirection{pos, dir}] = true
					pos.y += dir.x
					pos.x -= dir.y
				}
			}

			score1 += area * len(inPerimeter)
			score2 += area * sides

		}
	}

	fmt.Println("ANSWER 1:", score1)
	fmt.Println("ANSWER 2:", score2)

}

func main() {
	grid := parse()
	solve1(grid)
}
