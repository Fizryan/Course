package main

import "fmt"

func main() {
	type PemainBola struct {
		Nama   string
		Gol    int
		Assist int
	}

	pemain := []PemainBola{
		{"", 0, 0},
		{"", 0, 0},
		{"", 0, 0},
	}

	for i := range pemain {
		fmt.Scan(&pemain[i].Nama, &pemain[i].Gol, &pemain[i].Assist)
		i += 1
	}

	topGol := 0
	topAssist := 0
	var namaTopGol, namaTopAssist string
	for _, p := range pemain {
		if p.Gol > topGol {
			topGol = p.Gol
			namaTopGol = p.Nama
		}
	}
	for _, p := range pemain {
		if p.Assist > topAssist {
			topAssist = p.Assist
			namaTopAssist = p.Nama
		}
	}

	fmt.Println(namaTopGol)
	fmt.Println(namaTopAssist)
}
