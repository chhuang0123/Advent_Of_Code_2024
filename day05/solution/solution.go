package solution

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ParseUpdate(input string) []string {
	return strings.Split(strings.Trim(input, "\n"), ",")
}

func ValidateUpdate(rules []string, input []string) (int, error) {
	orderCount := len(input)
	for i := 0; i < orderCount-1; i++ {
		for j := i + 1; j < orderCount; j++ {
			rule := fmt.Sprintf("%s|%s", input[i], input[j])

			if !slices.Contains(rules, rule) {
				return 0, nil
			}
		}
	}

	return strconv.Atoi(input[orderCount/2])
}
