package main

import "fmt"

const nMax int = 100

type belanja struct {
	nama  string
	harga int
}

type keranjang [nMax]belanja

func inputBelanja(arrKeranjang *keranjang, n *int) {
	/*I.S. Terdefinisi array arrKeranjang dan bilangan bulat n*/
	/*F.S. Array arrKeranjang terisi dengan data yang telah diinputkan*/
	var nama string
	var harga int
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&nama, &harga)
		(*arrKeranjang)[i].nama = nama
		(*arrKeranjang)[i].harga = harga
	}
}

func tampilData(arrKeranjang *keranjang, n *int) {
	/*I.S. Terdefinisi array arrKeranjang dan bilangan bulat n*/
	/*F.S. Menampilkan daftar barang yang dibeli beserta total harga dari barang-barang tersebut*/
	var totalHarga, a int
	a = 0
	for i := 0; i < *n; i++ {
		a += 1
		fmt.Printf("%d. %s\n", a, (*arrKeranjang)[i].nama)
		totalHarga += (*arrKeranjang)[i].harga
	}
	fmt.Printf("Total belanja :  %d\n", totalHarga)
}

func main() {
	var arrKeranjang keranjang
	var n int
	inputBelanja(&arrKeranjang, &n)
	tampilData(&arrKeranjang, &n)
}
