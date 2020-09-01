package main

import "fmt"

func palindrom(word string) bool {
	length := len(word) - 1
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrom("0011100"))
}
