package main

import (
	"fmt"
)

func ganjil(N int) {
	if N <= 0 {
		return
	}

	if N%2 != 0 {
		ganjil(N - 2)

	} else if N%2 == 0 {
		ganjil(N - 1)
	}

	if N%2 != 0 {
		fmt.Printf("%d ", N)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	ganjil(N)
}
