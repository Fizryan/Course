package main

import "fmt"

func segitigaRecursive(baris int, kolom int, n int) {
	if kolom <= n {
		if n-kolom >= baris {
			fmt.Print(" ")
		} else {
			fmt.Print("*")
		}
		segitigaRecursive(baris, kolom+1, n)

	} else if baris < n {
		fmt.Println()
		segitigaRecursive(baris+1, 1, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	segitigaRecursive(1, 1, n)
}
