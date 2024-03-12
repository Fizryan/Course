package main

import (
	"fmt"
)

// Procedure untuk memperbesar gambar (zoom-in)
func zoomIn(w, h, scale int) (int, int) {
	return w * scale, h * scale
}

// Procedure untuk memperkecil gambar (zoom-out)
func zoomOut(w, h, scale int) (int, int) {
	return w / scale, h / scale
}

// Procedure untuk membaca masukan
func readInput() (int, int, rune, int) {
	var w, h, scale int
	var operation rune // Menggunakan rune untuk menerima karakter tunggal

	// Membaca masukan
	fmt.Scanf("%d %d\n", &w, &h)
	fmt.Scanf("%c %d\n", &operation, &scale)

	return w, h, operation, scale
}

// Procedure untuk memilih operasi dan menampilkan keluaran
func processAndPrint(w, h, scale int, operation rune) {
	// Memilih operasi berdasarkan simbol
	if operation == '+' {
		w, h = zoomIn(w, h, scale)
	} else if operation == '-' {
		w, h = zoomOut(w, h, scale)
	} else {
		fmt.Println("Simbol operasi tidak valid.")
		return
	}

	// Menampilkan keluaran
	fmt.Println(w, h)
}

func main() {
	w, h, operation, scale := readInput()
	processAndPrint(w, h, scale, operation)
}
