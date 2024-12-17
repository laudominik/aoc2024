package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	a      int
	b      int
	c      int
	pc     int
	output []int
}

type TEXT []int

func parse() (code TEXT) {
	data, _ := os.ReadFile("data/day17.txt")
	codeStr := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, dStr := range codeStr {
		d, _ := strconv.Atoi(dStr)
		code = append(code, d)
	}
	return code
}

func combo(cpu CPU, arg int) int {
	switch arg {
	case 0:
		return arg
	case 1:
		return arg
	case 2:
		return arg
	case 3:
		return arg
	case 4:
		return cpu.a
	case 5:
		return cpu.b
	case 6:
		return cpu.c
	}
	panic(0)
}

func execute(cpu CPU, op int, arg int) CPU {
	switch op {
	case 0: // adv
		cpu.a >>= combo(cpu, arg)
	case 1: // bxl
		cpu.b ^= arg
	case 2: // bst
		cpu.b = combo(cpu, arg) % 8
	case 3: // jnz
		if cpu.a != 0 {
			cpu.pc = arg
			return cpu
		}
	case 4: // bxc
		cpu.b ^= cpu.c
	case 5: // out
		cpu.output = append(cpu.output, combo(cpu, arg)%8)
	case 6: // bdv
		cpu.b = cpu.a >> combo(cpu, arg)
	case 7: // cdv
		cpu.c = cpu.a >> combo(cpu, arg)

	}
	cpu.pc += 2
	return cpu
}

func reset(a int) CPU {
	return CPU{a, 0, 0, 0, []int{}}
}

func solve1(cpu CPU, code TEXT) {
	fmt.Println("ANSWER 1:", run(cpu, code))
}

func run(cpu CPU, code TEXT) []int {
	for cpu.pc < len(code) {
		op := code[cpu.pc]
		arg := code[cpu.pc+1]
		cpu = execute(cpu, op, arg)
	}
	return cpu.output
}

func differ(arr []int, tgt []int) bool {
	for i, x := range arr {
		y := tgt[i]
		if x != y {
			return true
		}
	}
	return false
}

func reverse(s []int) (t []int) {
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return t
}

func solve2(code TEXT) {
	solution := 0
	codeRev := reverse(code)
	for j := range codeRev {
		solution *= 8
		for {
			output := run(reset(solution), code)
			if len(output) == j+1 && !differ(reverse(output), codeRev) {
				break
			}
			solution++
		}
	}

	fmt.Println("ANSWER 2: ", solution)
}

func main() {
	code := parse()
	solve1(reset(30878003), code)
	solve2(code)
}
