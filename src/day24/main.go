package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	arg0 string
	arg1 string
	gate string
	out  string
}

func parse() (state map[string]bool, design []Gate) {
	data1, _ := os.ReadFile("data/day24_part1.txt")
	data2, _ := os.ReadFile("data/day24_part2.txt")
	lines1 := strings.Split(strings.TrimSpace(string(data1)), "\n")
	lines2 := strings.Split(strings.TrimSpace(string(data2)), "\n")
	state = make(map[string]bool)
	for _, line := range lines1 {
		lineLexed := strings.Split(line, ":")
		key, val := lineLexed[0], lineLexed[1]
		if strings.TrimSpace(val) == "0" {
			state[key] = false
		} else {
			state[key] = true
		}
	}
	for _, line := range lines2 {
		fields := strings.Fields(line)
		design = append(design, Gate{fields[0], fields[2], fields[1], fields[4]})
	}
	return state, design
}

func forwardPass(state map[string]bool, des Gate) {
	arg0 := state[des.arg0]
	arg1 := state[des.arg1]
	switch des.gate {
	case "AND":
		state[des.out] = arg0 && arg1
	case "OR":
		state[des.out] = arg0 || arg1
	case "XOR":
		state[des.out] = (arg0 && (!arg1)) || ((!arg0) && arg1)
	}
}

func sww(z byte, state map[string]bool) int {
	var startswithZ [][2]string
	for k, s := range state {
		if k[0] == z && s {
			startswithZ = append(startswithZ, [2]string{k, "1"})
		}
		if k[0] == z && !s {
			startswithZ = append(startswithZ, [2]string{k, "0"})
		}
	}
	sort.Slice(startswithZ, func(i, j int) bool { return startswithZ[i][0] < startswithZ[j][0] })

	number := ""
	for _, v := range startswithZ {
		number = v[1] + number
	}

	numberNumber, _ := strconv.ParseInt(number, 2, 64)
	return int(numberNumber)
}

func solve1(state map[string]bool, design []Gate) {
	passed := make(map[int]bool)
	hasValue := make(map[string]bool)

	for k := range state {
		hasValue[k] = true
	}
	// n-pass
	for len(passed) != len(design) {
		for i, d := range design {
			if hasValue[d.arg0] && hasValue[d.arg1] && !passed[i] {
				forwardPass(state, d)
				hasValue[d.out] = true
				passed[i] = true
			}
		}
	}
	fmt.Println("ANSWER 1: ", sww('z', state))

	/* part 2 helpers */
	//X := sww('x', state)
	//Y := sww('y', state)
	//Z := sww('z', state)
	//fmt.Println(X, "+", Y, "=", Z)
	//fmt.Println(X+Y, "=", Z)
	//fmt.Println(strconv.FormatInt(int64(X+Y), 2))
	//fmt.Println(strconv.FormatInt(int64(Z), 2))
}

func solve2(design []Gate) {
	/* no code for this part, manual analysis of a circuit */
	outfile := "digraph {\n"
	for _, gate := range design {
		outfile += "\t" + gate.out + "-> {" + gate.arg0 + " " + gate.arg1 + "}"

		switch gate.gate {
		case "AND":
			outfile += "[color=blue]\n"
		case "OR":
			outfile += "[color=green]\n"
		case "XOR":
			outfile += "[color=red]\n"
		}
	}
	outfile += "}"

	dotFile := "data/day24.dot"
	os.WriteFile(dotFile, []byte(outfile), 0644)

	cmd := exec.Command("dot", "-Tpng", dotFile, "-o", "data/day24.png")
	cmd.Run()

	/* swaps... */
	// swap1: fkp <-> z06 (fixes bits z06, z07)
	// swap2: ngr <-> z11 (fixes bits z11, z12)
	// swap3: mfm <-> z31 (fixes bits z31, z32)
	// swap4: krj <-> bpt (fixes bits z38, z39, z40, z41)
	fmt.Println("ANSWER 2: bpt,fkp,krj,mfm,ngr,z06,z11,z31")
}

func main() {
	s, d := parse()
	solve1(s, d)
	solve2(d)
}
