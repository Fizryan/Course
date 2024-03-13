package main

import (
	"fmt"
)

func factorial(n float64) float64 {
	if n == 1 || n == 0 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}

func permutation(n, r float64) float64 {
	return factorial(n) / factorial(n-r)
}

func combination(n, r float64) float64 {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d float64
	fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)
	fmt.Printf("%.0f %.0f\n%.0f %.0f", permutation(a, c), combination(a, c), permutation(b, d), combination(b, d))

}
