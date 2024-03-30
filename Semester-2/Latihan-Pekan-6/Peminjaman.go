package main

import "fmt"

func main() {

	type Tanggal struct {
		Tanggal int
		Bulan   int
		Tahun   int
	}

	var pinjam, kembali Tanggal
	fmt.Scan(&pinjam.Tanggal, &pinjam.Bulan, &pinjam.Tahun)
	fmt.Scan(&kembali.Tanggal, &kembali.Bulan, &kembali.Tahun)

	totalPinjam := (kembali.Tahun - pinjam.Tahun) * 360
	totalPinjam += (kembali.Bulan - pinjam.Bulan) * 30
	totalPinjam += kembali.Tanggal - pinjam.Tanggal

	fmt.Println(totalPinjam)
}
