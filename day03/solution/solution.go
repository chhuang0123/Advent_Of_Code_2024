package solution

import (
	"fmt"
	"math"
	"regexp"
)

func ExtractMultiplication(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAllString(input, -1)
}

func SumOfMultiplication(allMul []string) int {
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

func ExtractMultiplicationWithCondition(input string) []string {
	var result []string

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	allMul := re.FindAllStringIndex(input, -1)

	re = regexp.MustCompile(`do\(\)`)
	allDo := re.FindAllStringIndex(input, -1)

	re = regexp.MustCompile(`don't\(\)`)
	allDont := re.FindAllStringIndex(input, -1)

	var allInvalid [][]int
	for _, dont := range allDont {
		allInvalid = append(allInvalid, []int{dont[0], findEnd(dont[0], allDo)})
	}

	for _, mul := range allMul {
		valid := true
		for _, invalid := range allInvalid {
			if mul[0] > invalid[0] && mul[0] < invalid[1] {
				valid = false
				break
			}
		}

		if valid {
			//fmt.Printf("%v %v\n", input[mul[0]:mul[1]], mul)
			result = append(result, input[mul[0]:mul[1]])
		}
	}

	return result
}

func findEnd(start int, all [][]int) int {
	end := math.MaxInt
	for _, value := range all {
		if value[0] > start {
			end = value[0]
			break
		}
	}

	return end
}
