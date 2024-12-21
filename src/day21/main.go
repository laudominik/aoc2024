package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+

	| 0 | A |
	+---+---+
*/
var keypad = map[string][2]int{
	"7": {0, 0}, "8": {0, 1}, "9": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"1": {2, 0}, "2": {2, 1}, "3": {2, 2},
	"0": {3, 1}, "A": {3, 2},
}

/*
	+---+---+
	| ^ | A |

+---+---+---+
| < | v | > |
+---+---+---+
*/
var dirpad = map[string][2]int{
	"^": {0, 1}, "A": {0, 2},
	"<": {1, 0}, "v": {1, 1}, ">": {1, 2},
}

func parse() []string {
	data, _ := os.ReadFile("data/day21.txt")
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

type Triplet struct {
	s     Pos
	t     Pos
	depth int
}

type Pos [2]int

func oobKeypad(p Pos) bool {
	return p[0] == 3 && p[1] == 0
}

func oobDirpad(p Pos) bool {
	return p[0] == 0 && p[1] == 0
}

func translate(s, t Pos, keypad bool) (out string) {
	dirChar := "<v^>"
	dirDir := [][2]int{
		{0, -1},
		{1, 0},
		{-1, 0},
		{0, 1},
	}

	dx := t[1] - s[1]
	dy := t[0] - s[0]
	dirIx := 0

	current := s
	for !(dx == 0 && dy == 0) {
		ch := dirChar[dirIx]
		dir := dirDir[dirIx]
		dirIx = (dirIx + 1) % 4

		move := 0
		var next Pos

		if dir[1] == 0 {
			// up down
			move = dy / dir[0]
			if move <= 0 {
				continue
			}
			next = Pos{current[0] + move*dir[0], current[1]}
		} else {
			// left right
			move = dx / dir[1]
			if move <= 0 {
				continue
			}
			next = Pos{current[0], current[1] + move*dir[1]}
		}

		if (keypad && oobKeypad(next)) || (!keypad && oobDirpad(next)) {
			continue
		}

		current = next
		for i := 0; i < move; i++ {
			out += string(ch)
		}
		dx = t[1] - current[1]
		dy = t[0] - current[0]
	}
	return out
}

func getMovesCount(s, t Pos, depth int, memory map[Triplet]int) int {
	sample := Triplet{s, t, depth}
	v, ok := memory[sample]
	if ok {
		return v
	}

	path := translate(s, t, false) + "A"
	if depth == 1 {
		memory[sample] = len(path)
		return memory[sample]
	}
	pathx := "A" + path
	out := 0
	for i := 1; i < len(pathx); i++ {
		s1 := dirpad[string(pathx[i-1])]
		t1 := dirpad[string(pathx[i])]
		out += getMovesCount(s1, t1, depth-1, memory)
	}
	memory[sample] = out
	return out
}

func solve(paths []string, dirpadNo int) (score int) {
	memory := make(map[Triplet]int)

	for _, path := range paths {
		translated := ""
		scoreLocal := 0

		pathx := "A" + path
		for i := 1; i < len(pathx); i++ {
			s := keypad[string(pathx[i-1])]
			t := keypad[string(pathx[i])]
			translated += translate(s, t, true) + "A"
		}
		fmt.Println(translated)
		translated = "A" + translated // start with "A"
		for i := 1; i < len(translated); i++ {
			s := dirpad[string(translated[i-1])]
			t := dirpad[string(translated[i])]
			scoreLocal += getMovesCount(s, t, dirpadNo, memory)
		}

		numstr := ""
		for _, c := range path {
			if unicode.IsDigit(c) {
				numstr += string(c)
			}
		}
		num, _ := strconv.Atoi(numstr)
		score += scoreLocal * num
	}

	return score
}

func main() {
	parsed := parse()
	answer1 := solve(parsed, 2)
	answer2 := solve(parsed, 25)
	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
