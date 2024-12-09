package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Triplet struct {
	id       int
	size     int
	position int
}

func parse() (fs []int) {
	data, _ := os.ReadFile("data/day9.txt")
	s := string(data)

	for i, c := range s {
		a, _ := strconv.Atoi(string(c))
		for j := 0; j < a; j++ {
			if i%2 == 0 {
				fs = append(fs, i/2)
			} else {
				fs = append(fs, -1)
			}
		}
	}
	return fs
}

func fs_checksum(fs []int) (checksum int) {
	for i, v := range fs {
		if v == -1 {
			continue
		}
		checksum += v * i
	}
	return checksum
}

func solve1(fs []int) {
	free_space_ix := 0
	for i := len(fs) - 1; i >= 0; i-- {
		if fs[i] == -1 {
			continue
		}
		for fs[free_space_ix] != -1 {
			free_space_ix++
		}

		if free_space_ix > i {
			break
		}

		fs[free_space_ix] = fs[i]
		fs[i] = -1
	}

	fmt.Println("ANSWER 1: ", fs_checksum(fs))
}

func make_free_map(fs []int) (free_map map[int][]int) {
	free_map = make(map[int][]int)
	for i := 0; i < len(fs); {
		current_len := 0
		for i < len(fs) && fs[i] == -1 {
			current_len++
			i++
		}
		free_map[current_len] = append(free_map[current_len], i-current_len)

		if i == len(fs) {
			break
		}

		for i < len(fs) && fs[i] != -1 {
			i++
		}
	}
	return free_map
}

func solve2(fs []int) {
	var file_queue []Triplet // id -> size

	for i := 0; i < len(fs); {
		for i < len(fs) && fs[i] == -1 {
			i++
		}
		if i == len(fs) {
			break
		}
		current_len := 0
		tag := fs[i]
		for i < len(fs) && fs[i] == tag {
			i++
			current_len++
		}
		file_queue = append(file_queue, Triplet{tag, current_len, i - current_len})
	}

	for i := len(file_queue) - 1; i >= 0; i-- {
		free_map := make_free_map(fs)
		file := file_queue[i]
		_ = free_map
		min := math.MaxInt
		for k, v := range free_map {
			if k < file.size {
				continue
			}
			if v[0] < min {
				min = v[0]
			}
		}
		if min == math.MaxInt || min > file.position {
			continue
		}
		for j := 0; j < file.size; j++ {
			fs[min+j] = file.id
			fs[file.position+j] = -1
		}
	}
	fmt.Println("\nANSWER 2: ", fs_checksum(fs))
}

func main() {
	fs1 := parse()
	solve1(fs1)
	fs2 := parse()
	solve2(fs2)
}
