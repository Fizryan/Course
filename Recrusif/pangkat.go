package main

import (
	"fmt"
)

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	result := pangkat(x, y-1)
	return x * result
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	result := pangkat(x, y)
	fmt.Printf("%d", result)
}
