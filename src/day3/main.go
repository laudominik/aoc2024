package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse() string {
	data, _ := os.ReadFile("data/day3.txt")
	return string(data)
}

func solve1(s string) {
	score := 0
	reMul, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	reDigit, _ := regexp.Compile(`\d+`)

	matches := reMul.FindAllString(s, -1)

	for _, match := range matches {
		nums := reDigit.FindAllString(match, -1)
		d1, _ := strconv.Atoi(nums[0])
		d2, _ := strconv.Atoi(nums[1])
		score += d1 * d2
	}

	fmt.Println("ANSWER 1: ", score)
}

func solve2(s string) {
	score := 0
	enabled := true
	reMul, _ := regexp.Compile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	reDigit, _ := regexp.Compile(`\d+`)

	matches := reMul.FindAllString(s, -1)

	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}

		nums := reDigit.FindAllString(match, -1)
		d1, _ := strconv.Atoi(nums[0])
		d2, _ := strconv.Atoi(nums[1])
		if enabled {
			score += d1 * d2
		}
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	s := parse()
	solve1(s)
	solve2(s)
}
