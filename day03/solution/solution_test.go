package solution

import (
	"slices"
	"testing"
)

func TestExtractMultiplication(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := ExtractMultiplication(input)
	want := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v, but want %v", got, want)
	}
}

func TestSumOfMultiplication(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := SumOfMultiplication(input)
	want := 161

	if got != want {
		t.Errorf("got: %v, but want %v", got, want)
	}
}
