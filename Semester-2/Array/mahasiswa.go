package main

import (
	"fmt"
)

// Kapasitas Array 1000
const maxWisudawan = 1000

type Wisudawan struct {
	Nama     string
	Nim      string
	Eprt     int
	Semester int
	Ipk      float64
}

var wisudawanList [maxWisudawan]Wisudawan
var count int

// Procedure Array Wisudawan
func inputWisudawan() {
	for {
		var wisudawan Wisudawan
		fmt.Print("Nama Wisudawan: ")
		fmt.Scan(&wisudawan.Nama)
		if wisudawan.Nama == "none" {
			break
		}
		fmt.Print("Masukan Nim, Eprt, Semester, Ipk: ")
		fmt.Scan(&wisudawan.Nim, &wisudawan.Eprt, &wisudawan.Semester, &wisudawan.Ipk)
		wisudawanList[count] = wisudawan
		count++
	}
}

// Function Mencari Eprt Tertinggi
func cariMaxEprt() int {
	maxEprt := 0
	for i := 0; i < count; i++ {
		if wisudawanList[i].Eprt > maxEprt {
			maxEprt = wisudawanList[i].Eprt
		}
	}
	return maxEprt
}

// Function Mencari Ipk Terendah
func cariMinIpk() float64 {
	if count == 0 {
		return 0
	}
	minIpk := wisudawanList[0].Ipk
	for i := 1; i < count; i++ {
		if wisudawanList[i].Ipk < minIpk {
			minIpk = wisudawanList[i].Ipk
		}
	}
	return minIpk
}

// Function Mencari Rata-Rata Semester
func rataRataSemester() float64 {
	if count == 0 {
		return 0
	}
	totalSemester := 0
	for i := 0; i < count; i++ {
		totalSemester += wisudawanList[i].Semester
	}
	return float64(totalSemester) / float64(count)
}

func main() {
	inputWisudawan()

	if count > 0 {
		fmt.Printf("Eprt tertinggi: %d \n", cariMaxEprt())
		fmt.Printf("Ipk terendah: %4.f \n", cariMinIpk())
		fmt.Printf("Rata-rata semester lulusan: %4.f \n", rataRataSemester())
	} else {
		fmt.Println("Tidak ada data wisudawan.")
	}
}
