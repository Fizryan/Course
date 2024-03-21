package main

import "fmt"

func faktorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func kombinasi(n, r int) int {
	if n == r || r == 0 {
		return 1
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	fmt.Println(kombinasi(n, r))
}
