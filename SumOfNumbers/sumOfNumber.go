package SumOfNumbers

// Given two integers a and b, which can be positive or negative,
// find the sum of all the integers between and including them and return it.
// If the two numbers are equal return a or b.

func GetSum(a, b int) int {
	if a > b {
		a, b = b, a
	}
	var ans int
	for i := a; i <= b; i++ {
		ans += i
	}
	return ans
}
