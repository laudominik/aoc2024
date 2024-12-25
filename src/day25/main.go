package main

import (
	"fmt"
	"os"
	"strings"
)

const BLOCK_HEIGHT = 7
const BLOCK_WIDTH = 5

type Block [BLOCK_WIDTH]int

func parse() (keys []Block, doors []Block) {
	data, _ := os.ReadFile("data/day25.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := 0; i < len(lines); i += BLOCK_HEIGHT + 1 {
		var kd Block
		isDoor := lines[i][0] == '#' || lines[i][1] == '#' || lines[i][2] == '#' || lines[i][3] == '#' || lines[i][4] == '#'
		for j := 0; j < BLOCK_HEIGHT; j++ {
			line := lines[i+j]
			if line[0] == '#' {
				kd[0]++
			}
			if line[1] == '#' {
				kd[1]++
			}
			if line[2] == '#' {
				kd[2]++
			}
			if line[3] == '#' {
				kd[3]++
			}
			if line[4] == '#' {
				kd[4]++
			}
		}
		if isDoor {
			doors = append(doors, kd)
		} else {
			keys = append(keys, kd)
		}
	}

	return keys, doors
}

func solve(keys []Block, doors []Block) {
	score := 0
	for _, key := range keys {
		for _, door := range doors {
			if key[0]+door[0] <= 7 &&
				key[1]+door[1] <= 7 &&
				key[2]+door[2] <= 7 &&
				key[3]+door[3] <= 7 &&
				key[4]+door[4] <= 7 {
				score++
			}
		}
	}
	fmt.Println("ANSWER: ", score)
}

func main() {
	keys, doors := parse()
	solve(keys, doors)
}
