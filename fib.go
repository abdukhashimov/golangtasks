package main

import "fmt"

func fib(number int) int {
	firstStep := 0
	secondSetp := 1
	var output = 0
	for i := 0; i <= number; i++ {
		output = firstStep + secondSetp
		firstStep = secondSetp
		secondSetp = output
	}
	return output
}

func main() {
	fmt.Println(fib(3))
}
