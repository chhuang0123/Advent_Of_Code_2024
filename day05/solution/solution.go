package solution

import (
	"fmt"
	"slices"
)

func ValidateUpdate(rules []string, input []string) bool {
	orderCount := len(input)
	for i := 0; i < orderCount-1; i++ {
		for j := i + 1; j < orderCount; j++ {
			rule := fmt.Sprintf("%s|%s", input[i], input[j])

			if !slices.Contains(rules, rule) {
				return false
			}
		}
	}

	return true
}
