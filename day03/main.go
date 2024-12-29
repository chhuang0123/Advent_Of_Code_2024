package main

import (
	"bufio"
	"day03/solution"
	"fmt"
	"os"
)

func main() {
	filename := "input"
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	total := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		total += solution.SumOfMultiplication(line)
	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d", total)
}
