package main

import (
	"bufio"
	"day06/solution"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input"
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	var guardMap [][]string
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		guardMap = append(guardMap, strings.Split(strings.TrimRight(line, "\n"), ""))
	}

	err = f.Close()
	if err != nil {
		return
	}

	part1Count := 0
	finalMap := solution.GetFinalMap(guardMap)
	for i := 0; i < len(finalMap); i++ {
		for j := 0; j < len(finalMap[0]); j++ {
			if finalMap[i][j] == "X" {
				part1Count++
			}
		}
	}

	fmt.Printf("part 1: %d\n", part1Count)
}
