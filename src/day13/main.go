package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

func parse() (as, bs, outs []Pair) {
	data, _ := os.ReadFile("data/day13.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	reDigit, _ := regexp.Compile(`\d+`)
	for i := 0; i < len(lines); i++ {
		if i%4 == 3 {
			continue
		}
		line := lines[i]
		match := reDigit.FindAllString(line, -1)
		x, _ := strconv.Atoi(match[0])
		y, _ := strconv.Atoi(match[1])
		switch i % 4 {
		case 0:
			as = append(as, Pair{x, y})
		case 1:
			bs = append(bs, Pair{x, y})
		case 2:
			outs = append(outs, Pair{x, y})
		}
	}
	return as, bs, outs
}

func solve(as, bs, outs []Pair, mod bool) (score int) {
	for i := 0; i < len(as); i++ {
		a := as[i]
		b := bs[i]
		o := outs[i]
		if mod {
			o.x += 10000000000000
			o.y += 10000000000000
		}

		// solution to
		// a.x * c_a + b.x * c_b = o.x
		// a.y * c_b + b.y * c_b = o.y
		// apparently each equation has one solution
		// (not necessarily a pair of integers)
		// so no linear programming needed
		D := a.x*b.y - a.y*b.x
		D_x := o.x*b.y - o.y*b.x
		D_y := a.x*o.y - a.y*o.x

		if D == 0 || D_y%D != 0 || D_x%D != 0 {
			continue
		}

		score += 3*(D_x/D) + (D_y / D)
	}
	return score
}

func main() {
	as, bs, outs := parse()
	answer1 := solve(as, bs, outs, false)
	answer2 := solve(as, bs, outs, true)

	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
