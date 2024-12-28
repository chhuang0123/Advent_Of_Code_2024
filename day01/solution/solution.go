package solution

import (
	"math"
	"sort"
)

func Part1Solution(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	var total float64 = 0
	for i := range a {
		result := math.Abs(float64(a[i] - b[i]))
		total = total + result
	}

	return int(total)
}

func Part2Solution(a []int, b []int) int {
	var valueCount = make(map[int]int)
	for _, value := range b {
		valueCount[value]++
	}

	total := 0
	for _, value := range a {
		total += value * valueCount[value]
	}

	return total
}
