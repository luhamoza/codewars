package SumOfNumbers_test

import (
	"github.com/luhamoza/codewars/SumOfNumbers"
	"testing"
)

func TestSumBetween(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 5, 15},    // Sum of integers from 1 to 5 is 1+2+3+4+5 = 15
		{-3, 2, -3},   // Sum of integers from -3 to 2 is -3+-2+-1+0+1+2 = -3
		{0, 0, 0},     // a and b are equal, so the sum is the value itself (0)
		{-5, -1, -15}, // Sum of integers from -5 to -1 is -5+-4+-3+-2+-1 = -15
	}

	for _, tc := range testCases {
		result := SumOfNumbers.GetSum(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("For range %d to %d, expected %d, but got %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
