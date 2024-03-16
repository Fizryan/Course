package main

import (
	"fmt"
)

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlahGanjil(bil1, bil2))
}

func jumlahGanjil(a, b int) int {
	sum := 0
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	return sum
}
