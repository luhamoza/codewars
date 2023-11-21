package FindOdd_test

import (
	"github.com/luhamoza/codewars/FindOdd"
	"testing"
)

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{7}, 7},
		{[]int{0}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 1, 0, 1, 0}, 0},
		{[]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}, 5},
	}

	for _, tc := range testCases {
		result := FindOdd.FindOdd(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}
