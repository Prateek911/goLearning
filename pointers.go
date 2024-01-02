package main

import "fmt"

func main() {
	var num int = 42
	var incCounterNum = num + 1
	var ptr *int
	var copyNum = num
	ptr = &num

	fmt.Println("num :", num)
	fmt.Println("copyNum :", copyNum)
	fmt.Println("Address:", &num)
	fmt.Println("Value through pointer:", *ptr)
	fmt.Println("incCounterNum :", incCounterNum)

}
