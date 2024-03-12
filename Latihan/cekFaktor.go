package main

import "fmt"

func cekFaktor(a, b int) {
	if b%a == 0 {
		fmt.Printf("Ya, %d faktor %d\n", a, b)
	} else {
		fmt.Printf("Tidak, %d bukan faktor %d\n", a, b)
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	cekFaktor(a, b)
}
