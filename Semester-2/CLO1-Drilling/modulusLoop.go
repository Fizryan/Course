package main

import "fmt"

func modulusLoop(n, m int) int {
	// Menggunakan loop untuk menghitung sisa hasil pembagian n dengan m
	for n >= m {
		n -= m
	}
	return n
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(modulusLoop(n, m))
}
