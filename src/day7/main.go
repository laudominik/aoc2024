package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() (targets []int, components [][]int) {
	data, _ := os.ReadFile("data/day7.txt")
	lines := strings.Split(strings.TrimSpace((string(data))), "\n")

	for _, line := range lines {
		split := strings.Split(strings.TrimSpace((string(line))), ":")
		target, _ := strconv.Atoi(split[0])
		targets = append(targets, target)
		fields := strings.Fields(split[1])
		var componentRow []int
		for _, field := range fields {
			a, _ := strconv.Atoi(field)
			componentRow = append(componentRow, a)
		}
		components = append(components, componentRow)
	}

	return targets, components
}

func isSolvable(currentValue, target, position int, components []int) bool {
	if position >= len(components) {
		return currentValue == target
	}

	return isSolvable(currentValue+components[position], target, position+1, components) ||
		isSolvable(currentValue*components[position], target, position+1, components)
}

func solve1(targets []int, components [][]int) {
	score := 0

	for i, target := range targets {
		if isSolvable(0, target, 0, components[i]) {
			score += target
		}
	}

	fmt.Println("ANSWER 1: ", score)
}

func isSolvableWithConcat(currentValue, target, position int, components []int) bool {

	if position >= len(components) {
		return currentValue == target
	}
	positionValue := components[position]

	strA := strconv.Itoa(currentValue)
	strB := strconv.Itoa(positionValue)
	concated, _ := strconv.Atoi(strA + strB)
	multipled := currentValue * positionValue
	added := currentValue + positionValue

	return isSolvableWithConcat(multipled, target, position+1, components) ||
		isSolvableWithConcat(added, target, position+1, components) ||
		isSolvableWithConcat(concated, target, position+1, components)
}

func solve2(targets []int, components [][]int) {
	score := 0

	for i, target := range targets {
		if isSolvableWithConcat(0, target, 0, components[i]) {
			score += target
		}
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	targets, components := parse()
	solve1(targets, components)
	solve2(targets, components)
}
