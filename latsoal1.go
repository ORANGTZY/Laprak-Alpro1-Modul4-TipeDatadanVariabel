package main

import "fmt"

func main() {
	var totalBelanja, diskon, totalAkhir int
	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukkan Diskon (%): ")
	fmt.Scan(&diskon)

	// Menghitung total harga setelah diskon
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)

	fmt.Println("Total belanja setelah diskon:", totalAkhir)
}
