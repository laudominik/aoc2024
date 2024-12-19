package main

import (
	"fmt"
	"os"
	"strings"
)

func parse() (patterns []string, lines []string) {
	part1, _ := os.ReadFile("data/day19_part1.txt")
	part2, _ := os.ReadFile("data/day19_part2.txt")
	patterns = strings.Split(strings.TrimSpace(string(part1)), ", ")
	lines = strings.Split(strings.TrimSpace(string(part2)), "\n")
	return patterns, lines
}

// DP approach
// dp[i] means s[0:i] can be constructed from patterns
func possibilities(s string, patterns []string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			if len(pattern) <= i && s[i-len(pattern):i] == pattern {
				dp[i] += dp[i-len(pattern)]
			}
		}
	}
	return dp[n]
}

func solve1(patterns []string, lines []string) {
	score := 0
	for _, line := range lines {
		if possibilities(line, patterns) > 0 {
			score++
		}
	}
	fmt.Println("ANSWER 1: ", score)
}

func solve2(patterns []string, lines []string) {
	score := 0
	for _, line := range lines {
		score += possibilities(line, patterns)
	}
	fmt.Println("ANSWER 2: ", score)
}

func main() {
	patterns, lines := parse()
	solve1(patterns, lines)
	solve2(patterns, lines)
}
