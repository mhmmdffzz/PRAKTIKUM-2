package main

import (
	"fmt"
)

// Fungsi untuk mengonversi suhu dari Celsius ke Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (9.0/5.0)*celsius + 32
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid, pastikan memasukkan angka positif.")
		return
	}

	temperatures := make([]float64, N)

	// Membaca suhu dalam Celsius
	fmt.Println("Masukkan suhu dalam Celsius:")
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&temperatures[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan memasukkan angka.")
			return
		}
	}

	// Mengonversi ke Fahrenheit dan mencetak hasil
	fmt.Println("Suhu dalam Fahrenheit:")
	for _, temp := range temperatures {
		fmt.Printf("%.2f\n", celsiusToFahrenheit(temp))
	}
}
