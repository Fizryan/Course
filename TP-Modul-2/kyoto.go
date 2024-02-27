package main

import (
	"fmt"
)

// Struct untuk merepresentasikan samurai
type Samurai struct {
	kekuatan  int
	kecepatan int
}

func main() {
	var n, kekuatan, kecepatan int
	var lawan []Samurai

	// Membaca input banyaknya samurai yang harus dihadapi Kenshin
	fmt.Scanln(&n)

	// Membaca kekuatan dan kecepatan Kenshin
	fmt.Scanln(&kekuatan, &kecepatan)

	// Membaca kekuatan dan kecepatan masing-masing lawan
	for i := 0; i < n; i++ {
		var power, speed int
		fmt.Scanln(&power, &speed)
		lawan = append(lawan, Samurai{kekuatan: power, kecepatan: speed})
	}

	// Proses pertarungan
	lawanYangDikalahkan := 0
	for i := 0; i < n; i++ {
		if lawan[i].kekuatan >= kekuatan && lawan[i].kecepatan >= kecepatan {
			continue
		} else if lawan[i].kekuatan >= kekuatan {
			kecepatan -= 6
		} else if lawan[i].kecepatan >= kecepatan {
			kekuatan -= 6
		} else {
			if kekuatan > kecepatan {
				kecepatan -= 6
			} else if kecepatan > kekuatan {
				kekuatan -= 6
			} else {
				kecepatan -= 6
			}
		}
		if kekuatan >= 3 && kecepatan >= 3 {
			lawanYangDikalahkan++
		} else {
			break
		}
	}

	// Output hasil pertarungan
	fmt.Println(lawanYangDikalahkan, kekuatan, kecepatan)
}
