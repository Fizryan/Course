package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(barisan(N))
}

func barisan(N int) int {
	if N == 1 || N == 2 || N == 3 {
		return 1
	}
	return barisan(N-1) + barisan(N-2) + barisan(N-3)
}
