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

func fibRec(number int) int {
	if number == 0 {
		return 1
	}

	if number == 1 {
		return 2
	}
	return fibRec(number-1) + fibRec(number-2)
}

func main() {
	fmt.Println(fibRec(8))
	fmt.Println(fib(8))
}
