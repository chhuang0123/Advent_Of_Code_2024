package main

import (
	"bufio"
	"day01/solution"
	"fmt"
	"os"
)

func main() {
	var (
		a []int
		b []int
	)

	filename := "input"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		var (
			left  int
			right int
		)

		_, err = fmt.Sscanf(line, "%d %d", &left, &right)
		if err != nil {
			fmt.Println(err)
			continue
		}

		a = append(a, left)
		b = append(b, right)
	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d\n", solution.Part1Solution(a, b))
	fmt.Printf("part 2: %d\n", solution.Part2Solution(a, b))
}
