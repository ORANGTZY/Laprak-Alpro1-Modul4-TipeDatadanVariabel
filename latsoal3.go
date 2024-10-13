package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	// Input koordinat titik A, B, dan C
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&cx, &cy)

	// Menghitung panjang sisi AB, BC, dan CA menggunakan teorema Pythagoras
	AB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	BC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	CA := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	// Menentukan sisi terpanjang
	sisiTerpanjang := math.Max(AB, math.Max(BC, CA))

	// Output hasil dengan 2 angka di belakang koma
	fmt.Printf("Sisi terpanjang dari segitiga tersebut adalah: %.2f\n", sisiTerpanjang)
}
