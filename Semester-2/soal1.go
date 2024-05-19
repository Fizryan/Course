package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol                                                                           [NMAX]int
	totPertandingan, totGol, totKebobolan, totMenang, totKalah, totDraw, totPoint int
}

func main() {
	var timE, timF tim
	bacaHasil(&timE, &timF)
	fmt.Println("Statistik Tim E")
	cetakHasil(timE)
	fmt.Println("\nStatistik Tim F")
	cetakHasil(timE)
}

func bacaHasil(tE, tF *tim) {
	/*
		IS: Tim E (tE) dan tim F (tF) terdefinisi sembarang
		Proses: Membaca skor pertandingan berupa jumlah gol tim E dan tim F.
				Mengisi field-field sesuai skor kedua tim. Field-fieldnya:
				1. total pertandingan
				2. gol setiap pertandingan
				3. total gol
				4. total kebobolan
				5. total point: point 3 untuk menang, point 1 untuk draw
				6. total menang: Menang, jika gol lebih besar dari gol lawan
				7. total draw: Draw, jika gol sama dengan gol lawan
				8. total kalah: Kalah, jika gol lebih kecil dari gol lawan
				Pembacaan skor berakhir jika kedua skor bernilai negatif.
		FS: Data tim E (tE) dan data tim F (tF) berisi nilai
	*/
	var golE, golF, i int
	i = 0
	fmt.Scan(&golE, &golF)
	for golE >= 0 && golF >= 0 && i < NMAX {
		tE.totPertandingan += 1
		tE.gol[i] = golE
		tF.gol[i] = golF
		tE.totGol += golE
		tE.totKebobolan += golF
		if golE > golF {
			tE.totMenang += 1
			tE.totPoint += 3
		}
		if golE < golF {
			tE.totKalah += 1
		}
		if golE == golF {
			tE.totDraw += 1
			tE.totPoint += 1
		}
		fmt.Scan(&golE, &golF)
		i++
	}
}

func cetakHasil(t tim) {
	/*
		IS: Data tim (t) terdefinisi
		FS: tercetak di layar statistik pertandingan tim (t) dengan format:

		Total pertandingan: <total pertandingan>
		Gol tiap pertandingan: <gol1 gol2 ... goln>
		Total menang: <total kemenangan>
		Total draw: <total draw>
		Total kalah: <total kalah>
		Total gol: <total gol>
		Total kebobolan: <total kebobolan>
		Total point: <total point>
	*/
	fmt.Printf("Total pertandingan: %d \n", t.totPertandingan)
	fmt.Print("Gol tiap pertandingan: ")
	for i := 0; i < t.totPertandingan; i++ {
		fmt.Printf("%d \n", t.gol[i])
	}
	fmt.Printf("Total menang: %d \n", t.totMenang)
	fmt.Printf("Total draw: %d \n", t.totDraw)
	fmt.Printf("Total kalah: %d \n", t.totKalah)
	fmt.Printf("Total gol: %d \n", t.totGol)
	fmt.Printf("Total kebobolan: %d \n", t.totKebobolan)
	fmt.Printf("Total point: %d \n", t.totPoint)
}
