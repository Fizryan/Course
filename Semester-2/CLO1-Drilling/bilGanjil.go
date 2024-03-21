package main

import (
	"fmt"
)

func hitungJumlah(b1, b2 int) int {
	/* function mengembalikan jumlah dua bilangan, jika kedua bilangan itu ganjil
	   Jika tidak ganjil keduanya, kembalikan bilangan 0
	*/
	if b1%2 != 0 && b2%2 != 0 {
		return b1 + b2
	} else {
		return 0
	}
}

func main() {
	var b1, b2 int
	fmt.Scan(&b1, &b2)
	fmt.Println(hitungJumlah(b1, b2))
}
