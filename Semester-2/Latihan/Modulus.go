package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(modulusLoop(n, m))
}

func modulusLoop(n, m int) int {
	var hasil int
	if n >= 0 && m > 0 {
		hasil = (n % m)
	} else {
		hasil = 0
	}
	return hasil
}
