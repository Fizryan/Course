package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Println("Masukan (P)Panjang sisi")
	fmt.Print("P sisi A: ")
	fmt.Scan(&a)
	fmt.Print("P sisi B: ")
	fmt.Scan(&b)
	fmt.Print("P sisi C: ")
	fmt.Scan(&c)

	result := jenis(a, b, c)

	fmt.Println(result)
}

func jenis(a, b, c int) string {
	if a == b && b == c {
		return "Segitiga Sama Sisi"
	} else if a == b || a == c || b == c {
		return "Segitiga Sama Kaki"
	} else {
		return "Segitiga Sembarang"
	}
}
