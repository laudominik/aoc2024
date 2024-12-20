package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	y int
	x int
}

func parse() []string {
	data, _ := os.ReadFile("data/day20.txt")
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func getStartEnd(grid []string) (startPos, endPos Pos) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				startPos = Pos{i, j}
			}
			if grid[i][j] == 'E' {
				endPos = Pos{i, j}
			}
		}
	}
	return startPos, endPos
}

func absI(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(grid []string, allowed int) int {
	startPos, endPos := getStartEnd(grid)
	path := getPath(grid, startPos, endPos)
	pathMap := make(map[Pos]int)

	score := 0
	bnd := 100

	var allowedArea []Pos
	for dy := -allowed; dy <= allowed; dy++ {
		for dx := -allowed; dx <= allowed; dx++ {
			if absI(dy)+absI(dx) > allowed {
				continue
			}
			allowedArea = append(allowedArea, Pos{dy, dx})
		}
	}

	for i, p := range path {
		pathMap[p] = i
	}

	for i, pt := range path {
		for _, d := range allowedArea {
			j, ok := pathMap[Pos{pt.y + d.y, pt.x + d.x}]
			if !ok {
				continue
			}
			if j < i || j == -1 {
				continue
			}
			if absI(i-j)-absI(d.y)-absI(d.x) >= bnd {
				score++
			}
		}
	}

	return score
}

func getPath(grid []string, start, end Pos) []Pos {
	current := start
	var path []Pos = []Pos{start}
	visited := make(map[Pos]bool)
	visited[start] = true
	for !(current.y == end.y && current.x == end.x) {
		nbs := neighbors(grid, current)
		for _, n := range nbs {
			if grid[n.y][n.x] == '#' || visited[n] {
				continue
			}
			visited[n] = true
			path = append(path, n)
			current = n
		}
	}
	return path
}

func neighbors(grid []string, pt Pos) []Pos {
	var nbs []Pos
	if pt.y > 0 {
		nbs = append(nbs, Pos{pt.y - 1, pt.x})
	}
	if pt.x > 0 {
		nbs = append(nbs, Pos{pt.y, pt.x - 1})
	}
	if pt.y < len(grid)-1 {
		nbs = append(nbs, Pos{pt.y + 1, pt.x})
	}
	if pt.x < len(grid[0])-1 {
		nbs = append(nbs, Pos{pt.y, pt.x + 1})
	}
	return nbs
}

func main() {
	grid := parse()
	answer1 := solve(grid, 2)
	answer2 := solve(grid, 20)
	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
