package main

import "fmt"

func faktor(i, N int) {
	if i <= N {
		if N%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(i+1, N)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	faktor(1, N)
}
