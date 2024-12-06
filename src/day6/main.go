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

type State struct {
	size     Vec2i
	barriers map[Vec2i]bool
	guard    Vec2i
}

func parse() (s State) {
	data, _ := os.ReadFile("data/day6.txt")

	lines := strings.Split(strings.TrimSpace((string(data))), "\n")
	s.barriers = make(map[Vec2i]bool)
	s.size = Vec2i{len(lines), len(lines[0])}

	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				s.guard = Vec2i{y, x}
			} else if c == '#' {
				s.barriers[Vec2i{y, x}] = true
			}
		}
	}
	return s
}

func solve1(s State) (visited map[Vec2i]bool) {
	visited = make(map[Vec2i]bool)
	pos := s.guard
	dir := Vec2i{-1, 0}

	for {
		visited[pos] = true
		for s.barriers[Vec2i{pos.y + dir.y, pos.x + dir.x}] {
			dir.x, dir.y = -dir.y, dir.x
		}
		pos.y += dir.y
		pos.x += dir.x
		if pos.y >= s.size.y || pos.x >= s.size.x || pos.y < 0 || pos.x < 0 {
			break
		}
	}

	return visited
}

func contains(list []Vec2i, a Vec2i) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func checkForLoop(s State) bool {
	visited := make(map[Vec2i]bool)
	dirsVisited := make(map[Vec2i][]Vec2i)
	pos := s.guard
	dir := Vec2i{-1, 0}

	for {
		visited[pos] = true
		dirsVisited[pos] = append(dirsVisited[pos], dir)

		for s.barriers[Vec2i{pos.y + dir.y, pos.x + dir.x}] {
			dir.x, dir.y = -dir.y, dir.x
		}
		pos.y += dir.y
		pos.x += dir.x
		if pos.y >= s.size.y || pos.x >= s.size.x || pos.y < 0 || pos.x < 0 {
			return false
		}

		if visited[pos] && contains(dirsVisited[pos], dir) {
			return true
		}
	}
}

func solve2(s State) {
	visited := solve1(s)
	visited[s.guard] = false
	score := 0

	for v, _ := range visited {
		s.barriers[v] = true
		if checkForLoop(s) {
			score++
		}
		s.barriers[v] = false
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	s := parse()
	visited := solve1(s)
	fmt.Println("ANSWER 1: ", len(visited))
	solve2(s)
}
