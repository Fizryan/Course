package main

import (
	"fmt"
)

func main() {
	var enemy, power, speed int
	fmt.Print("Input Jumlah Musuh : ")
	fmt.Scan(&enemy)
	fmt.Print("Input Power dan Speed : ")
	fmt.Scan(&power, &speed)

	KO := 0
	for i := 0; i < enemy; i++ {
		var kekuatan, kelincahan int
		fmt.Print("Input Power dan Speed : ")
		fmt.Scan(&kekuatan, &kelincahan)
		if (kekuatan > power && kelincahan > speed) || (kekuatan == power && kelincahan == speed) {
			continue
		} else if kekuatan >= power {
			speed -= 6
		} else if kelincahan >= speed {
			power -= 6
		} else {
			if power > speed {
				speed -= 6
			} else {
				power -= 6
			}
		}
		if power >= 3 && speed >= 3 {
			KO++
		} else {
			break
		}
	}
	fmt.Println(KO, power, speed)
}
