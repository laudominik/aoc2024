package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
	y int
	x int
}

type Maze []string
type Memo [][]int

const INF = math.MaxInt / 2

func parse() (maze Maze) {
	data1, _ := os.ReadFile("data/day16.txt")
	maze = strings.Split(strings.TrimSpace(string(data1)), "\n")
	return maze
}

func (maze Maze) print() {
	for _, row := range maze {
		println(row)
	}
}

func solve(maze Maze) {
	start_y := 0
	start_x := 0
	end_y := 0
	end_x := 0
	var memo Memo
	bestPath := make(map[Pos]bool)

	for y, row := range maze {
		var memo_row []int
		for x, c := range row {
			if c == 'S' {
				start_y = y
				start_x = x
			}
			if c == 'E' {
				end_y = y
				end_x = x
			}
			memo_row = append(memo_row, INF)

		}
		memo = append(memo, memo_row)
	}
	fmt.Println(start_y, start_x)
	fillCostFx(maze, start_y, start_x, 0, 1, memo, 0)

	cost := memo[end_y][end_x]
	fmt.Println("ANSWER 1: ", cost)

	fillBestPath(cost, end_y, end_x, memo, bestPath)
	fmt.Println("ANSWER 2: ", len(bestPath))
}

func fillCostFx(maze Maze, y, x, dir_y, dir_x int, memory Memo, current int) {
	if y < 0 || x < 0 || y >= len(maze) || x >= len(maze[0]) || maze[y][x] == '#' || memory[y][x] < current {
		return
	}
	memory[y][x] = current
	if maze[y][x] == 'E' {
		return
	}
	fillCostFx(maze, y+dir_y, x+dir_x, dir_y, dir_x, memory, current+1)
	fillCostFx(maze, y-dir_x, x+dir_y, -dir_x, dir_y, memory, current+1001)
	fillCostFx(maze, y+dir_x, x-dir_y, dir_x, -dir_y, memory, current+1001)
}

func fillBestPath(cost, y, x int, memory Memo, bestPath map[Pos]bool) {
	costCurrent := memory[y][x]

	if cost < 0 || costCurrent != cost {
		return
	}

	bestPath[Pos{y, x}] = true
	fillBestPath(cost-1, y+1, x, memory, bestPath)
	fillBestPath(cost-1001, y+1, x, memory, bestPath)
	fillBestPath(cost-1001, y, x-1, memory, bestPath)
	fillBestPath(cost-1, y, x-1, memory, bestPath)
	fillBestPath(cost-1001, y, x+1, memory, bestPath)
	fillBestPath(cost-1, y, x+1, memory, bestPath)
	fillBestPath(cost-1001, y-1, x, memory, bestPath)
	fillBestPath(cost-1, y-1, x, memory, bestPath)

}

func main() {
	maze := parse()
	maze.print()
	solve(maze)
}
