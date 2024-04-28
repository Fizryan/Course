package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var nama string
	fmt.Scan(&nama)
	for nama != "none" {
		A[*n].nama = nama
		fmt.Scan(&A[*n].nomorPunggung, &A[*n].posisi, &A[*n].tinggi)
		*n += 1

		if *n == 11 {
			break
		}

		fmt.Scan(&nama)
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	var a pemain
	fmt.Println("Data Pemain:")
	for i := 0; i < n; i++ {
		a = A[i]
		fmt.Printf("%s %s %s %d\n", a.nama, a.nomorPunggung, a.posisi, a.tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var max int = 0
	var a pemain
	for i := 1; i < n; i++ {
		a = A[i]
		if a.tinggi > A[max].tinggi {
			max = i
		}
	}
	fmt.Printf("Pemain dengan badan tertinggi: %s\n", A[max].nama)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var min int = 0
	var a pemain
	for i := 1; i < n; i++ {
		a = A[i]
		if a.tinggi < A[min].tinggi {
			min = i
		}
	}
	fmt.Printf("Pemain dengan badan terendah: %s", A[min].nama)
}
