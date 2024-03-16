package main

import (
	"fmt"
)

func kali(n int) int {
	/*{Terdefinisi sebuah bilangan bulat positif n.
	  Fungsi mengembalikan hasil penjumlahan 1, 2, 3, ...., n. }*/
	if n == 1 {
		return n
	} else {
		return n * kali(n-1)
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(kali(n))
}
