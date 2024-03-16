package main

import (
	"fmt"
)

func main() {
	// Langkah 1: Masukkan nilai dalam satuan detik
	var t int
	fmt.Print("Masukkan nilai dalam satuan detik: ")
	fmt.Scan(&t)

	// Langkah 2: Hitung jumlah jam (jt)
	jt := t / 3600

	// Langkah 3: Hitung sisa detik (s)
	s := t % 3600

	// Langkah 4: Hitung jumlah menit (mt)
	mt := s / 60

	// Langkah 5: Hitung sisa detik (dt)
	dt := s % 60

	// Langkah 6: Tampilkan hasil denominasi jam, menit, dan detik
	fmt.Printf("%d jam %d menit dan %d detik\n", jt, mt, dt)
}
