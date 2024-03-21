package main

import "fmt"

func pangkat(a, b int) int {
	if b == 1 {
		return b
	} else {
		return a * pangkat(a, b-1)
	}
}

func main() {
	var a, N int
	fmt.Scan(&a)
	fmt.Scan(&N)
	fmt.Print(pangkat(a, N))
}
