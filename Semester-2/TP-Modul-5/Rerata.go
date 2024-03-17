package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	cetakRerata(N, 1, 0, 0)
}

func cetakRerata(N, i, jumlah int, rata int) {
	if i > N {
		return
	}
	jumlah += i
	rata = float64(jumlah) / float64(i)
	fmt.Println(jumlah, rata)
	cetakRerata(N, i+1, jumlah, rata)
}
