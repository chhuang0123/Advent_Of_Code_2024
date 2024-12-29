package solution

import (
	"fmt"
	"regexp"
)

func ExtractMultiplication(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAllString(input, -1)
}

func SumOfMultiplication(input string) int {
	allMul := ExtractMultiplication(input)
	total := 0

	for _, value := range allMul {
		var left, right int

		if _, err := fmt.Sscanf(value, "mul(%d,%d)", &left, &right); err != nil {
			fmt.Println(err)
			continue
		}
		total += left * right
	}

	return total
}
