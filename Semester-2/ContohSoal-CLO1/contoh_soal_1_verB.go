package main

import "fmt"

func cetakGanjilverB(n, i int) int {
	if i == n {
		if i%2 != 0 {
			fmt.Print(" ", n)
			return n
		} else {
			return n - 1
		}
	}
	if i%2 != 0 {
		fmt.Print(" ", i)
		return cetakGanjilverB(n, i+1)
	}

	return cetakGanjilverB(n, i+1)
}

func main() {
	var N int
	fmt.Scan(&N)
	cetakGanjilverB(N, 1)
}
