package solution

import (
	"slices"
	"testing"
)

func TestValidateUpdate(t *testing.T) {
	rules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	t.Run("true cases", func(t *testing.T) {
		input := []string{"75", "47", "61", "53", "29"}
		got := ValidateUpdate(rules, input)

		if got != true {
			t.Errorf("got: %v, but want %v", got, true)
		}
	})

	t.Run("true cases", func(t *testing.T) {
		input := []string{"97", "61", "53", "29", "13"}
		got := ValidateUpdate(rules, input)

		if got != true {
			t.Errorf("got: %v, but want %v", got, true)
		}
	})

	t.Run("true cases", func(t *testing.T) {
		input := []string{"75", "29", "13"}
		got := ValidateUpdate(rules, input)

		if got != true {
			t.Errorf("got: %v, but want %v", got, true)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"75", "97", "47", "61", "53"}
		got := ValidateUpdate(rules, input)

		if got != false {
			t.Errorf("got: %v, but want %v", got, false)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"61", "13", "29"}
		got := ValidateUpdate(rules, input)

		if got != false {
			t.Errorf("got: %v, but want %v", got, false)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"97", "13", "75", "29", "47"}
		got := ValidateUpdate(rules, input)

		if got != false {
			t.Errorf("got: %v, but want %v", got, false)
		}
	})
}

func TestGetCorrectOrderingUpdate(t *testing.T) {
	rules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	t.Run("case1", func(t *testing.T) {
		input := []string{"75", "97", "47", "61", "53"}
		got := GetCorrectOrderingUpdate(rules, input)
		want := []string{"97", "75", "47", "61", "53"}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("case2", func(t *testing.T) {
		input := []string{"61", "13", "29"}
		got := GetCorrectOrderingUpdate(rules, input)
		want := []string{"61", "29", "13"}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("case3", func(t *testing.T) {
		input := []string{"97", "13", "75", "29", "47"}
		got := GetCorrectOrderingUpdate(rules, input)
		want := []string{"97", "75", "47", "29", "13"}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})
}
