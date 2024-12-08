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

type set map[Vec2i]bool
type freq_map map[rune]set

func parse() (antennas freq_map, bnds Vec2i) {
	data, _ := os.ReadFile("data/day8.txt")
	lines := strings.Split(strings.TrimSpace((string(data))), "\n")
	bnds = Vec2i{len(lines), len(lines[0])}

	antennas = make(freq_map)

	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				continue
			}
			_, ok := antennas[c]
			if !ok {
				antennas[c] = make(set)
			}
			antennas[c][Vec2i{y, x}] = true
		}
	}

	return antennas, bnds
}

func solve1(antennas freq_map, bnds Vec2i) {
	score := 0
	ps := make(set)

	for _, ants := range antennas {
		for pt1 := range ants {
			for pt2 := range ants {
				if pt1.x == pt2.x && pt1.y == pt2.y {
					continue
				}
				ps[Vec2i{pt1.y*2 - pt2.y, pt1.x*2 - pt2.x}] = true
				ps[Vec2i{pt2.y*2 - pt1.y, pt2.x*2 - pt1.x}] = true
				if (pt2.x-pt1.x)%3 == 0 && (pt2.y-pt1.y)%3 == 0 {
					ps[Vec2i{(2*pt1.y + pt2.y) / 3, (2*pt1.x + pt2.x) / 3}] = true
					ps[Vec2i{(2*pt2.y + pt1.y) / 3, (2*pt2.x + pt1.x) / 3}] = true
				}
			}
		}
	}
	for pt := range ps {
		if pt.y >= 0 && pt.y < bnds.y && pt.x >= 0 && pt.x < bnds.x {
			score++
		}
	}

	fmt.Println("ANSWER 1: ", score)
}

func solve2(antennas freq_map, bnds Vec2i) {
	ps := make(set)

	for _, ants := range antennas {
		for pt1 := range ants {
			for pt2 := range ants {
				if pt1.x == pt2.x && pt1.y == pt2.y {
					continue
				}
				dx := pt2.x - pt1.x
				dy := pt2.y - pt1.y
				pt := pt2
				for {
					if pt.y < 0 || pt.y >= bnds.y || pt.x < 0 || pt.x >= bnds.x {
						break
					}
					ps[pt] = true
					pt.x += dx
					pt.y += dy
				}
				pt = pt2
				for {
					if pt.y < 0 || pt.y >= bnds.y || pt.x < 0 || pt.x >= bnds.x {
						break
					}
					ps[pt] = true
					pt.x -= dx
					pt.y -= dy
				}

			}
		}
	}

	fmt.Println("ANSWER 2: ", len(ps))
}

func main() {
	antennas, bnds := parse()
	solve1(antennas, bnds)
	solve2(antennas, bnds)
}
