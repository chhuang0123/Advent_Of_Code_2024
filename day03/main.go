package main

import (
	"bufio"
	"day03/solution"
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

	var (
		part1Total = 0
		part2Total = 0
	)

	var allLines strings.Builder
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		allLines.WriteString(line)
	}

	err = f.Close()
	if err != nil {
		return
	}

	allMul := solution.ExtractMultiplication(allLines.String())
	part1Total += solution.SumOfMultiplication(allMul)

	allMul = solution.ExtractMultiplicationWithCondition(allLines.String())
	part2Total += solution.SumOfMultiplication(allMul)

	fmt.Printf("part 1: %d\n", part1Total)
	fmt.Printf("part 2: %d\n", part2Total)
}
