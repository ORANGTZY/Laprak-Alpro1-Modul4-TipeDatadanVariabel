package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (dalam meter): ")
	fmt.Scan(&tinggiBadan)

	// Menghitung berat badan berdasarkan BMI dan tinggi badan
	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("Berat badan anda adalah: %.2f kg\n", beratBadan)
}
