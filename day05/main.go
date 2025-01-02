package main

import (
	"bufio"
	"day05/solution"
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

	rules := []string{}
	part1Count := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Index(line, "|") != -1 {
			rules = append(rules, strings.Trim(line, "\n"))
		} else if strings.Index(line, ",") != -1 {
			result, _ := solution.ValidateUpdate(rules, solution.ParseUpdate(line))
			part1Count += result
		}

	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d\n", part1Count)
}
