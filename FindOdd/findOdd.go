package FindOdd

import (
	"fmt"
)

func FindOdd(seq []int) int {
	answer := make(map[int]int)
	for _, s := range seq {
		answer[s]++
	}
	fmt.Println(answer)
	var oddElement int
	for _, s := range seq {
		if answer[s]%2 != 0 {
			oddElement = s
			break
		}
	}
	return oddElement
}
