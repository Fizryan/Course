package main

import (
	"fmt"
)

func main() {

	var profitBulanIni, profitBulanSebelumnya float64

	fmt.Print("Masukkan keuntungan bulan ini: ")
	fmt.Scan(&profitBulanIni)
	fmt.Print("Masukkan keuntungan bulan sebelumnya: ")
	fmt.Scan(&profitBulanSebelumnya)

	selisih := profitBulanIni - profitBulanSebelumnya

	var status string
	if selisih > 0 {
		status = "peningkatan"
	} else if selisih < 0 {
		status = "penurunan"
	} else {
		status = "tidak ada perubahan"
	}

	fmt.Printf("%s keuntungan sebesar %.2f.\n", status, selisih)
}
