package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() (arr [][]int) {
	data, _ := os.ReadFile("data/day2.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)
		var lineArr []int
		for _, f := range fields {
			val, _ := strconv.Atoi(f)
			lineArr = append(lineArr, val)
		}
		arr = append(arr, lineArr)
	}
	return arr
}

func breakIndex(arr []int, descending bool) int {
	for i := 0; i < len(arr)-1; i++ {
		var diff int

		if descending {
			diff = arr[i] - arr[i+1]
		} else {
			diff = arr[i+1] - arr[i]
		}
		if !(1 <= diff && diff <= 3) {
			return i
		}
	}
	return len(arr)
}

/* O(nm) */
func solve1(arr [][]int) {
	score := 0

	for _, lineArr := range arr {
		if breakIndex(lineArr, true) == len(lineArr) || breakIndex(lineArr, false) == len(lineArr) {
			score++
		}
	}

	fmt.Println("ANSWER 1: ", score)
}

func copyRemove(slice []int, ix int) (cpyArr []int) {
	for j, v := range slice {
		if ix == j {
			continue
		}
		cpyArr = append(cpyArr, v)
	}
	return cpyArr
}

/* O(n m^2) */
func solve2(arr [][]int) {
	score := 0
	_ = score
	for _, lineArr := range arr {
		var N = len(lineArr)
		var ix0 = breakIndex(lineArr, false)
		var ix1 = breakIndex(lineArr, true)
		if ix0 == N || ix1 == N {
			score++
			continue
		}

		for i, _ := range lineArr {
			cpyArr := copyRemove(lineArr, i)

			ix0 = breakIndex(cpyArr, false)
			ix1 = breakIndex(cpyArr, true)
			if ix0 == N-1 || ix1 == N-1 {
				score++
				break
			}
		}
	}

	fmt.Println("ANSWER 2: ", score)
}

func main() {
	arr := parse()
	solve1(arr)
	solve2(arr)
}
