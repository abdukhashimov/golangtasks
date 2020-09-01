package main

import "fmt"

func oddEvenSum(nums ...int) (int, int) {
	var oddSum, evenSum int = 0, 0

	for _, element := range nums {
		if element%2 == 0 {
			evenSum += element
		} else {
			oddSum += element
		}
	}
	return oddSum, evenSum
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(oddEvenSum(nums...))
}
