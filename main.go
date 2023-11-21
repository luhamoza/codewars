package main

import (
	"fmt"
	"github.com/luhamoza/codewars/BuildTower"
	"github.com/luhamoza/codewars/FindOdd"
	"github.com/luhamoza/codewars/PrinterError"
	"github.com/luhamoza/codewars/SumOfNumbers"
)

func main() {
	FindOdd.FindOdd([]int{0})
	SumOfNumbers.GetSum(0, 0)
	PrinterError.PrintError("")
	fmt.Println(BuildTower.TowerBuilder(6))
}
