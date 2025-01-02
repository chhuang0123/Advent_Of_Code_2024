package solution

import (
	"slices"
	"testing"
)

func TestParseUpdate(t *testing.T) {
	input := "75,47,61,53,29"
	got := ParseUpdate(input)
	want := []string{"75", "47", "61", "53", "29"}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, but want %v", got, want)
	}
}

func TestTestValidateUpdate(t *testing.T) {
	rules := []string{"47|53",
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
		got, _ := ValidateUpdate(rules, input)
		want := 61

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("true cases", func(t *testing.T) {
		input := []string{"97", "61", "53", "29", "13"}
		got, _ := ValidateUpdate(rules, input)
		want := 53

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("true cases", func(t *testing.T) {
		input := []string{"75", "29", "13"}
		got, _ := ValidateUpdate(rules, input)
		want := 29

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"75", "97", "47", "61", "53"}
		got, _ := ValidateUpdate(rules, input)
		want := 0

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"61", "13", "29"}
		got, _ := ValidateUpdate(rules, input)
		want := 0

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})

	t.Run("false cases", func(t *testing.T) {
		input := []string{"97", "13", "75", "29", "47"}
		got, _ := ValidateUpdate(rules, input)
		want := 0

		if got != want {
			t.Errorf("got: %v, but want %v", got, want)
		}
	})
}
