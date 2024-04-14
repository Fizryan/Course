package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	var in int

	*n = 0
	fmt.Scan(&in)

	for in != 0 && *n < 10 {
		A[*n] = in
		*n++
		fmt.Scan(&in)
	}
}

func cetak(A tabInt, n int) {
	var i int

	if n > 0 {
		for i = 0; i < n; i++ {
			if A[i] > 0 {
				fmt.Print(A[i], " ")
			}
		}
	}
	fmt.Println()

}

func jumlah(A tabInt, n int) int {
	var i, jumlah int

	jumlah = 0
	for i = 0; i < n; i++ {
		if A[i] < 0 {
			A[i] *= -1
		}
		jumlah += A[i]
	}
	return jumlah
}

func rata_rata(A tabInt, n int) float64 {
	var jum, mean float64

	jum = float64(jumlah(A, n))
	mean = jum / float64(n)
	return mean
}
