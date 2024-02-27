package main

import "fmt"

func main() {

    var adik, kakak int
    fmt.Scan(&adik)
    fmt.Scan(&kakak)
	adikWin := false

	fmt.Println("Masukan Nomor!! ")
	fmt.Scan(&adik, &kakak)
	fmt.Println("Nomor Adik ", adik)
	fmt.Println("Nomor Kakak ", kakak)
	if adik%2 == 0 && kakak%2 != 0 || adik%2 != 0 && kakak%2 == 0 {
		adikWin = true
		fmt.Println("Adik Menang!!! ", adik, kakak, "status kemenangan adik: ", adikWin)
	} else {
		adikWin = false
		fmt.Println("Tidak ada yang menang ", adik, kakak, "status kemenangan adik: ", adikWin)
	}
}