package main

import (
	"fmt"
)

func main() {
	var s string
	var n int

	fmt.Scan(&s, &n)
	recPrint(n, s)
}

func recPrint(n int, s string) {
	/*I.S. Terdefinisi nilai bilangan bulat sebanyak n.
	  F.S.  menampilkan string sebanyak n kali*/

	if n == 1 {
		fmt.Println(s)
	} else {
		fmt.Println(s)
		recPrint(n-1, s)
	}
}
