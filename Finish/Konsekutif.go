package main

import (
	"fmt"
)

func isConsecutive(num uint) bool {
	current := int(num)%10 - (int(num)/10)%10
	for num > 10 {
		if current > 1 || current < -1 || current == 0 {
			return false
		} else {
			num /= 10
		}
	}
	return true
}

func main() {
	var x uint
	fmt.Printf("Enter a positive integer\n")
	fmt.Scan(&x)
	fmt.Printf("Consecutive : %t", isConsecutive(x))
}
