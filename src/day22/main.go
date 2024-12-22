package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RNG struct {
	secret int
}

func parse() (seeds []int) {
	data, _ := os.ReadFile("data/day22.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines {
		a, _ := strconv.Atoi(line)
		seeds = append(seeds, a)
	}
	return seeds
}

func makeRNG(seed int) RNG {
	rng := RNG{seed}
	return rng
}

func mix(a int, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func (rng *RNG) generate() int {
	secret := rng.secret
	secret = prune(mix(secret*64, secret))
	secret = prune(mix(secret/32, secret))
	secret = prune(mix(secret*2048, secret))
	rng.secret = secret
	return secret
}

func (rng *RNG) generateNth(N int) int {
	for i := 0; i < N; i++ {
		rng.generate()
	}
	return rng.secret
}

func solve1(seeds []int) (score int) {
	for _, seed := range seeds {
		rng := makeRNG(seed)
		score += rng.generateNth(2000)
	}
	return score
}

func sequenceReward(seed int, stats map[[4]int]int) {
	seen := make(map[[4]int]bool)
	sequence := [4]int{-1, -1, -1, -1}
	rng := makeRNG(seed)
	sequence[0] = rng.secret % 10
	sequence[1] = rng.generate() % 10
	sequence[2] = rng.generate() % 10
	sequence[3] = rng.generate() % 10

	for i := 3; i < 2000; i++ {
		gen := rng.generate() % 10
		diffSequence := [4]int{
			sequence[1] - sequence[0],
			sequence[2] - sequence[1],
			sequence[3] - sequence[2],
			gen - sequence[3],
		}
		sequence[0] = sequence[1]
		sequence[1] = sequence[2]
		sequence[2] = sequence[3]
		sequence[3] = gen

		if seen[diffSequence] {
			continue
		}
		seen[diffSequence] = true
		_, ok := stats[diffSequence]
		if ok {
			stats[diffSequence] += gen
		} else {
			stats[diffSequence] = gen
		}
	}
}

func solve2(seeds []int) (score int) {
	stats := make(map[[4]int]int)
	for _, seed := range seeds {
		sequenceReward(seed, stats)
	}
	mx := 0
	for _, v := range stats {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func main() {
	seeds := parse()
	ans1 := solve1(seeds)
	ans2 := solve2(seeds)
	fmt.Println("ANSWER 1: ", ans1)
	fmt.Println("ANSWER 2: ", ans2)
}
