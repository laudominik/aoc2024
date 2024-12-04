package main

import (
	"fmt"
	"os"
	"strings"
)

func parse() []string {
	data, _ := os.ReadFile("data/day4.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}

func omniMatch(ss []string, x, y, bnd_x, bnd_y int) int {
	/* match X M A S in all directions */
	matches := 0
	if x >= 3 && ss[y][x-3] == 'S' && ss[y][x-2] == 'A' && ss[y][x-1] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	if x <= bnd_x-4 && ss[y][x] == 'X' && ss[y][x+1] == 'M' && ss[y][x+2] == 'A' && ss[y][x+3] == 'S' {
		matches++
	}

	if y >= 3 && ss[y-3][x] == 'S' && ss[y-2][x] == 'A' && ss[y-1][x] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	if y <= bnd_y-4 && ss[y][x] == 'X' && ss[y+1][x] == 'M' && ss[y+2][x] == 'A' && ss[y+3][x] == 'S' {
		matches++
	}

	if x >= 3 && y >= 3 && ss[y-3][x-3] == 'S' && ss[y-2][x-2] == 'A' && ss[y-1][x-1] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	if x >= 3 && y <= bnd_y-4 && ss[y+3][x-3] == 'S' && ss[y+2][x-2] == 'A' && ss[y+1][x-1] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	if x <= bnd_x-4 && y >= 3 && ss[y-3][x+3] == 'S' && ss[y-2][x+2] == 'A' && ss[y-1][x+1] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	if x <= bnd_x-4 && y <= bnd_y-4 && ss[y+3][x+3] == 'S' && ss[y+2][x+2] == 'A' && ss[y+1][x+1] == 'M' && ss[y][x] == 'X' {
		matches++
	}

	return matches
}

func omniMatch2(ss []string, x, y int) int {
	/* matches X-MAS flipped whatever*/
	matches := 0
	/*
		M S
		 A
		M S
	*/
	if ss[y][x] == 'M' && ss[y][x+2] == 'S' &&
		ss[y+1][x+1] == 'A' &&
		ss[y+2][x] == 'M' && ss[y+2][x+2] == 'S' {
		matches++
	}

	/*
		M M
		 A
		S S
	*/
	if ss[y][x] == 'M' && ss[y][x+2] == 'M' &&
		ss[y+1][x+1] == 'A' &&
		ss[y+2][x] == 'S' && ss[y+2][x+2] == 'S' {
		matches++
	}

	/*
	  S M
	   A
	  S M
	*/
	if ss[y][x] == 'S' && ss[y][x+2] == 'M' &&
		ss[y+1][x+1] == 'A' &&
		ss[y+2][x] == 'S' && ss[y+2][x+2] == 'M' {
		matches++
	}

	/*
	  S	S
	   A
	  M M
	*/
	if ss[y][x] == 'S' && ss[y][x+2] == 'S' &&
		ss[y+1][x+1] == 'A' &&
		ss[y+2][x] == 'M' && ss[y+2][x+2] == 'M' {
		matches++
	}

	return matches
}

func solve1(ss []string) {
	score := 0
	bnd_y := len(ss)
	bnd_x := len(ss[0])
	println(bnd_y, bnd_x)

	for i, l := range ss {
		for j, c := range l {
			if c != 'X' {
				continue
			}

			score += omniMatch(ss, j, i, bnd_x, bnd_y)
		}
	}

	fmt.Println("ANSWER 1: ", score)
}

func solve2(ss []string) {
	score := 0
	for i := 0; i < len(ss)-2; i++ {
		for j := 0; j < len(ss[0])-2; j++ {
			score += omniMatch2(ss, j, i)
		}
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	ss := parse()
	solve1(ss)
	solve2(ss)
}
