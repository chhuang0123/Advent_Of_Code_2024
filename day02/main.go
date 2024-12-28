package main

import (
	"bufio"
	"day02/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	part1SafeCount, part2SafeCount := 0, 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		// convert from string to int
		var record []int
		for _, str := range strings.Fields(line) {
			if value, err := strconv.Atoi(str); err != nil {
				break
			} else {
				record = append(record, value)
			}
		}

		if solution.IsAllIncreasingOrDecreasing(record) {
			part1SafeCount++
			part2SafeCount++
		} else {
			if solution.IsAllIncreasingOrDecreasingWithErrorTolerance(record) {
				part2SafeCount++
			}
		}
	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d\n", part1SafeCount)
	fmt.Printf("part 2: %d\n", part2SafeCount)
}
