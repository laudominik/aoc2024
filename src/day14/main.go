package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

const bound_x = 101
const bound_y = 103
const bins_no = 5
const max_unchanged_iters = 10000 /* empyrical */

func parse() (ps, vs []Pair) {
	data, _ := os.ReadFile("data/day14.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	reDigit, _ := regexp.Compile(`[-]?\d+`)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		match := reDigit.FindAllString(line, -1)
		x, _ := strconv.Atoi(match[0])
		y, _ := strconv.Atoi(match[1])
		v_x, _ := strconv.Atoi(match[2])
		v_y, _ := strconv.Atoi(match[3])
		ps = append(ps, Pair{x, y})
		vs = append(vs, Pair{v_x, v_y})
	}
	return ps, vs
}

func solve1(ps, vs []Pair) int {
	iters := 100
	quad0 := 0
	quad1 := 0
	quad2 := 0
	quad3 := 0

	for i, p := range ps {
		v := vs[i]
		out_x := (v.x*iters + p.x) % bound_x
		out_y := (v.y*iters + p.y) % bound_y

		if out_x < 0 {
			out_x += bound_x
		}
		if out_y < 0 {
			out_y += bound_y
		}

		if out_x == bound_x/2 || out_y == bound_y/2 {
			continue
		}
		if out_x < bound_x/2 && out_y < bound_y/2 {
			quad0++
			continue
		}
		if out_x > bound_x/2 && out_y < bound_y/2 {
			quad1++
			continue
		}
		if out_x < bound_x/2 && out_y > bound_y/2 {
			quad2++
			continue
		}
		if out_x > bound_x/2 && out_y > bound_y/2 {
			quad3++
			continue
		}
		panic(0)
	}
	return quad0 * quad1 * quad2 * quad3
}

// wtf is "a picture of a Christmas tree"
// low-entropy histogram
func solve2(ps, vs []Pair) (iters int) {
	total := float64(len(ps))
	minEntropy := math.MaxFloat64
	minIter := 0
	itersSinceLastChange := 0

	for {
		iters++

		histogram := make([][]int, bins_no)
		for i := range histogram {
			histogram[i] = make([]int, bins_no)
		}

		for i, p := range ps {
			v := vs[i]
			p.x = (v.x + p.x) % bound_x
			p.y = (v.y + p.y) % bound_y

			if p.x < 0 {
				p.x += bound_x
			}
			if p.y < 0 {
				p.y += bound_y
			}

			histogram[p.x*bins_no/bound_x][p.y*bins_no/bound_y]++
			ps[i] = p
		}

		entropy := 0.0
		for _, row := range histogram {
			for _, v := range row {
				if v == 0 {
					continue
				}
				count := float64(v)
				p := count / total
				entropy -= p * math.Log2(p)
			}
		}
		if entropy < minEntropy {
			minEntropy = entropy
			itersSinceLastChange = 0
			minIter = iters
		}

		if itersSinceLastChange > max_unchanged_iters {
			break
		}
		itersSinceLastChange++
	}
	return minIter
}

func main() {
	ps, vs := parse()
	answer1 := solve1(ps, vs)
	answer2 := solve2(ps, vs)

	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
