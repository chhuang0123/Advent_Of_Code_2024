package main

import (
	"bufio"
	"day05/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input"
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	var rules []string
	part1Count, part2Count := 0, 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Index(line, "|") != -1 {
			rules = append(rules, strings.Trim(line, "\n"))
		} else if strings.Index(line, ",") != -1 {
			update := strings.Split(strings.Trim(line, "\n"), ",")
			if solution.ValidateUpdate(rules, update) {
				if center, err := strconv.Atoi(update[len(update)/2]); err == nil {
					part1Count += center
				}
			} else {
				newUpdate := solution.GetCorrectOrderingUpdate(rules, update)
				if center, err := strconv.Atoi(newUpdate[len(newUpdate)/2]); err == nil {
					part2Count += center
				}
			}
		}
	}

	err = f.Close()
	if err != nil {
		return
	}

	fmt.Printf("part 1: %d\n", part1Count)
	fmt.Printf("part 2: %d\n", part2Count)
}
