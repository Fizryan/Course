package main

import "fmt"

func banyakBilanganGanjil(N int) int {
	/*  fungsi mengembalikan  bilangan bulat yang menyatakan
	    banyaknya bilangan ganjil dari 1 hingga N */
	if N == 1 {
		return 1
	}
	if N%2 == 1 {
		return 1 + banyakBilanganGanjil(N-2)
	}
	return banyakBilanganGanjil(N - 1)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(banyakBilanganGanjil(N))
}
