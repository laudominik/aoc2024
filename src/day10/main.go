package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	y int
	x int
}

type stack []Position

func (s stack) Push(v Position) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, Position) {
	if len(s) == 0 {
		panic(0)
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

func parse() (grid [][]int, trailheads []Position) {
	data, _ := os.ReadFile("data/day10.txt")
	gridS := strings.Split(strings.TrimSpace(string(data)), "\n")

	for y, s := range gridS {
		var gridRow []int
		for x, c := range s {
			if c == '0' {
				trailheads = append(trailheads, Position{y, x})
			}
			val, _ := strconv.Atoi(string(c))
			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
	return grid, trailheads
}

func solve(grid [][]int, trailheads []Position, distinct bool) (score int) {
	height := len(grid)
	width := len(grid[0])

	for _, t := range trailheads {
		var st stack
		visited := make(map[Position]bool)
		st = st.Push(t)
		for len(st) != 0 {
			var a Position
			st, a = st.Pop()
			if visited[a] && !distinct {
				continue
			}
			visited[a] = true
			val := grid[a.y][a.x]
			if val == 9 {
				score++
			}
			if a.y != 0 && grid[a.y-1][a.x]-val == 1 {
				st = st.Push(Position{a.y - 1, a.x})
			}
			if a.x != 0 && grid[a.y][a.x-1]-val == 1 {
				st = st.Push(Position{a.y, a.x - 1})
			}
			if a.y != height-1 && grid[a.y+1][a.x]-val == 1 {
				st = st.Push(Position{a.y + 1, a.x})
			}
			if a.x != width-1 && grid[a.y][a.x+1]-val == 1 {
				st = st.Push(Position{a.y, a.x + 1})
			}
		}
	}

	return score

}

func main() {
	grid, trailheads := parse()
	score1 := solve(grid, trailheads, false)
	score2 := solve(grid, trailheads, true)

	fmt.Println("ANSWER 1: ", score1)
	fmt.Println("ANSWER 2: ", score2)
}
