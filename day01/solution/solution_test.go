package solution

import (
	"testing"
)

func TestPart1Solution(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	got := Part1Solution(a, b)
	var want = 11

	if got != want {
		t.Errorf("got: %v, but want %v", got, want)
	}
}

func TestPart2Solution(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	got := Part2Solution(a, b)
	var want = 31

	if got != want {
		t.Errorf("got: %v, but want %v", got, want)
	}
}
