package main

import "fmt"

func cetakGanjil(n int) {
	if n == 1 {
		fmt.Print(n)
		return
	}
	cetakGanjil(n - 1)

	if n%2 != 0 {
		fmt.Print(" ", n)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	cetakGanjil(N)
}
