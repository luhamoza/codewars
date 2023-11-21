package PrinterError_test

import (
	"github.com/luhamoza/codewars/PrinterError"
	"testing"
)

func TestPrinterError(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"aaabbbbhaijjjm", "0/14"},         // No errors, all characters are in the range 'a' to 'm'
		{"aaaxbbbbyyhwawiwjjjwwm", "8/22"}, // 8 errors: 'x', 'y', 'h', 'w', 'w', 'w', 'w', 'm'
		{"abcdefghi", "0/9"},               // All characters are errors as they are not in the range 'a' to 'm'
		{"abcdefghijklm", "0/13"},          // No errors, all characters are in the range 'a' to 'm'
		{"a", "0/1"},                       // No errors, only one character in the string
		{"abcdefghijklmno", "2/15"},        // 1 error: 'n'
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := PrinterError.PrintError(tc.input)
		if result != tc.expected {
			t.Errorf("For input %s, expected %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
