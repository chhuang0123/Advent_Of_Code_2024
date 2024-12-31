package solution

import (
	"testing"
)

func TestCheckVertical(t *testing.T) {
	var tests = []struct {
		input [][]string
		want  int
	}{
		{[][]string{{".", ".", "X", ".", ".", "."},
			{".", "S", "A", "M", "X", "."},
			{".", "A", ".", ".", "A", "."},
			{"X", "M", "A", "S", ".", "S"},
			{".", "X", ".", ".", ".", "."}},
			1},
	}

	for _, tt := range tests {
		t.Run("is there a vertical XMAS?", func(t *testing.T) {
			got := CheckVertical(tt.input, 4, 1)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckHorizontal(t *testing.T) {
	var tests = []struct {
		input [][]string
		want  int
	}{
		{[][]string{{".", ".", "X", ".", ".", "."},
			{".", "S", "A", "M", "X", "."},
			{".", "A", ".", ".", "A", "."},
			{"X", "M", "A", "S", ".", "S"},
			{".", "X", ".", ".", ".", "."}},
			1},
	}

	for _, tt := range tests {
		t.Run("is there a horizontal XMAS?", func(t *testing.T) {
			got := CheckHorizontal(tt.input, 1, 4)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
		t.Run("is there a horizontal XMAS?", func(t *testing.T) {
			got := CheckHorizontal(tt.input, 3, 0)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckDiagonal(t *testing.T) {
	var tests = []struct {
		input [][]string
		want  int
	}{
		{[][]string{{".", ".", "X", ".", ".", "."},
			{".", "S", "A", "M", "X", "."},
			{".", "A", ".", ".", "A", "."},
			{"X", "M", "A", "S", ".", "S"},
			{".", "X", ".", ".", ".", "."}},
			1},
	}

	for _, tt := range tests {
		t.Run("is there a diagonal XMAS?", func(t *testing.T) {
			got := CheckDiagonal(tt.input, 0, 2)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckAllXmas(t *testing.T) {
	var tests = []struct {
		input [][]string
		want  int
	}{
		{[][]string{
			{".", ".", "X", ".", ".", "."},
			{".", "S", "A", "M", "X", "."},
			{".", "A", ".", ".", "A", "."},
			{"X", "M", "A", "S", ".", "S"},
			{".", "X", ".", ".", ".", "."}},
			4},
		{[][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
			{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
			{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
			{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
			{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
			{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
			{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
			{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"}},
			18},
	}

	for _, tt := range tests {
		t.Run("get all XMAS", func(t *testing.T) {
			got := CheckAllXmas(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckAllCrossMas(t *testing.T) {
	var tests = []struct {
		input [][]string
		want  int
	}{
		{[][]string{
			{".", "M", ".", "S", ".", ".", ".", ".", ".", "."},
			{".", ".", "A", ".", ".", "M", "S", "M", "S", "."},
			{".", "M", ".", "S", ".", "M", "A", "A", ".", "."},
			{".", ".", "A", ".", "A", "S", "M", "S", "M", "."},
			{".", "M", ".", "S", ".", "M", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{"S", ".", "S", ".", "S", ".", "S", ".", "S", "."},
			{".", "A", ".", "A", ".", "A", ".", "A", ".", "."},
			{"M", ".", "M", ".", "M", ".", "M", ".", "M", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."}},
			9},
	}

	for _, tt := range tests {
		t.Run("get all X-MAS", func(t *testing.T) {
			got := CheckAllCrossMas(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}
