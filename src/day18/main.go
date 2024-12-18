package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Memo [][]int
type Set map[Pos]bool
type Pos struct {
	y int
	x int
}

const N = 71
const FIRST_BYTES = 1024
const INF = math.MaxInt / 2

func parse() (obstacles []Pos) {
	data, _ := os.ReadFile("data/day18.txt")
	lines := strings.Split(strings.TrimSpace((string(data))), "\n")
	for _, line := range lines {
		fields := strings.Split(strings.TrimSpace((string(line))), ",")
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])
		obstacles = append(obstacles, Pos{b, a})
	}
	return obstacles
}

func resetMemory() (memory Memo) {
	for i := 0; i < N; i++ {
		var row []int
		for j := 0; j < N; j++ {
			row = append(row, INF)
		}
		memory = append(memory, row)
	}
	return memory
}

func runWithFirstN(obstacles []Pos, n int) int {
	obstaclesMap := make(Set)
	for i := 0; i < n; i++ {
		obstacle := obstacles[i]
		obstaclesMap[obstacle] = true
	}
	memory := resetMemory()
	fillCostFx(obstaclesMap, 0, 0, memory, 0)
	return memory[N-1][N-1]
}

func solve(obstacles []Pos) {
	fmt.Println("ANSWER 1: ", runWithFirstN(obstacles, FIRST_BYTES))
	left := 0
	right := len(obstacles) - 1
	for {
		if left+1 == right {
			break
		}
		pivot := (right + left) / 2
		val := runWithFirstN(obstacles, pivot)
		if val == INF {
			right = pivot
		} else {
			left = pivot
		}
	}
	coord := obstacles[left]
	print("ANSWER 2: ", coord.x, ",", coord.y)
	println("")
}

func fillCostFx(obstacles Set, y, x int, memory Memo, current int) {
	if y < 0 || x < 0 || y >= N || x >= N || obstacles[Pos{y, x}] || memory[y][x] <= current {
		return
	}
	memory[y][x] = current
	if y == N-1 && x == N-1 {
		return
	}
	fillCostFx(obstacles, y+1, x, memory, current+1)
	fillCostFx(obstacles, y, x+1, memory, current+1)
	fillCostFx(obstacles, y-1, x, memory, current+1)
	fillCostFx(obstacles, y, x-1, memory, current+1)
}

func main() {
	obstacles := parse()
	solve(obstacles)
}
