package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	resultSum := add(5, 4)
	fmt.Println("Sum:", resultSum)
	resultDiff := diff(5, 4)
	fmt.Println("Diff:", resultDiff)
	resultProduct := multiply(5, 4)
	fmt.Println("Diff:", resultProduct)
}
