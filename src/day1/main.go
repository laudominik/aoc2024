package main

import (
    "fmt"
    "os"
	"strings"
    "strconv"
    "sort"
)

func parse() (arr1, arr2 []int) {
    data, _ := os.ReadFile("data/day1.txt")  
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")

    for _, line := range lines {
        fields := strings.Fields(line)
        a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])

        arr1 = append(arr1, a)
        arr2 = append(arr2, b)
    }
    return arr1, arr2
}

/* O(nlogn) - sorting */
func solve1(arr1, arr2 []int) {
    score := 0

    sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
    sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })
    for i := 0; i < len(arr1); i++ {
        diff := arr1[i] - arr2[i]
        if diff < 0 {
            diff *= -1
        }
        score += diff
	}
    fmt.Println("ANSWER 1: ", score)
}

/* O(nlogn) - sorting */
func solve2(arr1, arr2 []int) {
    score := 0
    sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
    sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

    it2 := 0
    for it1 := 0; it1 < len(arr1); it1++ {
        if it2 >= len(arr2) { 
            break
        }
        count := 0
        for it2 < len(arr2) && arr1[it1] > arr2[it2] {
            it2++
        }
        for it2 < len(arr2) && arr1[it1] == arr2[it2] {
            count++
            it2++
        }
        score += count * arr1[it1]
        count = 0
    }
    fmt.Println("ANSWER 2: ", score)
}

func main() {
    arr1, arr2 := parse();
    solve1(arr1, arr2)
    solve2(arr1, arr2)
}
