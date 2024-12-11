package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	number int
	depth  int
}

type memory map[Pair]int

func parse() (numbers []int) {
	data, _ := os.ReadFile("data/day11.txt")
	for _, field := range strings.Fields(string(data)) {
		var val, _ = strconv.Atoi(field)
		numbers = append(numbers, val)
	}
	return numbers
}

func blink(number int, depth int, mem memory) int {
	if depth == 0 {
		return 1
	}

	nd := Pair{number, depth}
	r, ok := mem[nd]
	if ok {
		return r
	}

	if number == 0 {
		mem[nd] = blink(1, depth-1, mem)
		return mem[nd]
	}

	numberAsStr := strconv.Itoa(number)

	if len(numberAsStr)%2 == 0 {
		split := len(numberAsStr) / 2
		number1, _ := strconv.Atoi(numberAsStr[:split])
		number2, _ := strconv.Atoi(numberAsStr[split:])
		mem[nd] = blink(number1, depth-1, mem) + blink(number2, depth-1, mem)
		return mem[Pair{number, depth}]

	}

	mem[nd] = blink(number*2024, depth-1, mem)
	return mem[nd]
}

func solve(numbers []int, depth int) (score int) {
	mem := make(memory)
	for _, number := range numbers {
		score += blink(number, depth, mem)
	}
	return score
}

func main() {
	numbers := parse()
	answer1 := solve(numbers, 25)
	answer2 := solve(numbers, 75)

	fmt.Println("ANSWER 1: ", answer1)
	fmt.Println("ANSWER 2: ", answer2)
}
