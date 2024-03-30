package main

import "fmt"

func main() {
	type Koordinat struct {
		X int
		Y int
	}
	var hartaKarun Koordinat
	fmt.Scan(&hartaKarun.X, &hartaKarun.Y)

	for i := 0; i < 3; i++ {
		var tebakan Koordinat
		fmt.Scan(&tebakan.X, &tebakan.Y)

		if tebakan.X == hartaKarun.X && tebakan.Y == hartaKarun.Y {
			fmt.Println("Yay Menang")
			return
		} else {
			fmt.Println("Coba Lagi")
		}
	}
}
