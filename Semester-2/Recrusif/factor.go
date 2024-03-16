package main

import (
	"fmt"
)

func factor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	factor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	factor(n, 1)

}
