package main

import (
	"bufio"
	"day04/solution"
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

	var wordSearch [][]string

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		wordSearch = append(wordSearch, strings.Split(strings.TrimRight(line, "\n"), ""))
	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d\n", solution.CheckAllXmas(wordSearch))
	fmt.Printf("part 2: %d\n", solution.CheckAllCrossMas(wordSearch))
}
