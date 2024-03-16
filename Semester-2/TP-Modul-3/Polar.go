package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64 {
	x := T * math.Pi / 180
	return R * math.Cos(x)
}

func panjangY(R, T float64) float64 {
	y := T * math.Pi / 180
	return R * math.Sin(y)
}

func main() {
	var R, T float64
	fmt.Scanf("%f %f", &R, &T)
	fmt.Printf("%.2f %.2f\n", panjangX(R, T), panjangY(R, T))
}
