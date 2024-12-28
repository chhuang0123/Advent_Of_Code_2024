package solution

import "testing"

func TestIsAllIncreasing(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, false},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		t.Run("is all integer are increasing at least 1 and at most 3", func(t *testing.T) {
			got := IsAllIncreasing(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsAllDecreasing(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, false},
	}

	for _, tt := range tests {
		t.Run("is all integer are decreasing at least 1 and at most 3", func(t *testing.T) {
			got := IsAllDecreasing(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsAllIncreasingOrDecreasing(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		t.Run("is all integer are increasing or decreasing at least 1 and at most 3", func(t *testing.T) {
			got := IsAllIncreasingOrDecreasing(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsAllIncreasingOrDecreasingWithErrorTolerance(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, tt := range tests {
		t.Run("is all integer are increasing or decreasing at least 1 and at most 3 and with an error tolerance ", func(t *testing.T) {
			got := IsAllIncreasingOrDecreasingWithErrorTolerance(tt.input)
			if got != tt.want {
				t.Errorf("%v: got: %v, but want: %v", tt.input, got, tt.want)
			}
		})
	}
}
