package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama                    string
	besar_bonus, masa_kerja  int
	gaji_pokok                     float64
}

func main() {
    // deklarasi variabel bertipe pegawai
	var pgw pegawai
	
	// baca data pengawai
	fmt.Scan(&pgw.nama, &pgw.gaji_pokok, &pgw.masa_kerja)
	
	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&pgw)
	
	// Cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %d\n", pgw.nama, pgw.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	if p.masa_kerja >= 10 {
		p.besar_bonus = int(p.gaji_pokok) * 15 / 10
	} else if p.masa_kerja >= 5 {
		p.besar_bonus = int(p.gaji_pokok) * 75 / 100
	} else {
		p.besar_bonus = int(p.gaji_pokok) * 5 / 10
	}
}