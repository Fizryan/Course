package main

import "fmt"

const nMax int = 1000

type tabInt [nMax]int

func isiArray(arr1, arr2 *tabInt, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr1, arr2 telah terisi sejumlah bilangan */
	for i := 0; i < n; i++ {
		fmt.Scan(&(*arr1)[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&(*arr2)[i])
	}
}

func hitungOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi or yang memenuhi kondisi pada array*/
	total := 0
	for i := 0; i < n; i++ {
		if arr1[i] == 1 || arr2[i] == 1 {
			total++
		}
	}
	return total
}

func hitungAND(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi logika and yang memenuhi kondisi pada array*/
	total := 0
	for i := 0; i < n; i++ {
		if arr1[i] == 1 && arr2[i] == 1 {
			total++
		}
	}
	return total
}

func hitungXOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi xor yang memenuhi kondisi pada array*/
	total := 0
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			total++
		}
	}
	return total
}

func cetak(arr1, arr2 tabInt, n int) {
	/* I.S. array arr1,arr2 berisi sejumlah n bilangan 1 dan 0
	F.S. Banyaknya hasil dari operasi logika OR, hasil dari operasi logika OR, banyaknya hasil dari operasi logika AND
	hasil dari operasi logika AND, banyaknya hasil dari operasi logika XOR,hasil dari operasi logika OR ditampilkan di layar */
	totalOR := hitungOR(arr1, arr2, n)
	fmt.Println("Total OR:", totalOR) // menampilkan total OR yang memenuhi
	fmt.Print("Hasil OR: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr1[i]|arr2[i], " ") // menampilkan hasil dari operasi logika OR
	}
	fmt.Println()

	totalAND := hitungAND(arr1, arr2, n)
	fmt.Println("Total AND:", totalAND) // menampilkan total operasi logika AND yang memenuhi
	fmt.Print("Hasil AND: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr1[i]&arr2[i], " ") // menampilkan hasil dari operasi logika AND
	}
	fmt.Println()

	totalXOR := hitungXOR(arr1, arr2, n)
	fmt.Println("Total XOR:", totalXOR) // menampilkan total operasi logika XOR yang memenuhi
	fmt.Print("Hasil XOR: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr1[i]^arr2[i], " ") // menampilkan hasil dari operasi logika XOR
	}
	fmt.Println()
}

func main() {
	var n int
	var arr1, arr2 tabInt

	fmt.Scan(&n)

	isiArray(&arr1, &arr2, n)
	cetak(arr1, arr2, n)
}
