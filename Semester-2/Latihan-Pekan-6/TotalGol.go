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

	totalGol := 0
	totalAssist := 0
	for _, p := range pemain {
		totalGol += p.Gol
		totalAssist += p.Assist
	}

	fmt.Println(totalGol)
	fmt.Println(totalAssist)
}
