package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	before int
	after  int
}

func parse() (rel map[Pair]bool, pages [][]int) {
	part1, _ := os.ReadFile("data/day5_part1.txt")
	part2, _ := os.ReadFile("data/day5_part2.txt")
	rel = make(map[Pair]bool)

	lines1 := strings.Split(strings.TrimSpace((string(part1))), "\n")
	lines2 := strings.Split(strings.TrimSpace((string(part2))), "\n")

	for _, line := range lines1 {
		fields := strings.Split(strings.TrimSpace((string(line))), "|")
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])

		rel[Pair{a, b}] = true
	}

	for _, line := range lines2 {
		fields := strings.Split(strings.TrimSpace((string(line))), ",")
		var page []int
		for _, f := range fields {
			c, _ := strconv.Atoi(f)
			page = append(page, c)
		}
		pages = append(pages, page)
	}

	return rel, pages
}

func solve1(rel map[Pair]bool, pages [][]int) {
	score := 0

	for _, page := range pages {
		correct := true
		for i, a := range page {
			for j := i + 1; j < len(page); j++ {
				b := page[j]
				if rel[Pair{b, a}] {
					correct = false
				}
			}
		}
		if correct {
			score += page[len(page)/2]
		}
	}

	fmt.Println("ANSWER 1: ", score)
}

func solve2(rel map[Pair]bool, pages [][]int) {
	score := 0

	for _, page := range pages {
		correct := true
		for i, a := range page {
			for j := i + 1; j < len(page); j++ {
				b := page[j]
				if rel[Pair{b, a}] {
					correct = false
				}
			}
		}
		if !correct {
			sort.Slice(page, func(i, j int) bool { return rel[Pair{page[i], page[j]}] })
			score += page[len(page)/2]
		}
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	rel, pages := parse()
	solve1(rel, pages)
	solve2(rel, pages)
}
