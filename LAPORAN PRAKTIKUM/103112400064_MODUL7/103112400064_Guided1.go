package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah tanaman: ")
	fmt.Scan(&n)

	var tinggiTanaman [500]float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan tinggi tanaman ke-%d (cm): ", i+1)
		fmt.Scan(&tinggiTanaman[i])
	}

	min, max := tinggiTanaman[0], tinggiTanaman[0]
	for i := 1; i < n; i++ {
		if tinggiTanaman[i] < min {
			min = tinggiTanaman[i]
		}
		if tinggiTanaman[i] > max {
			max = tinggiTanaman[i]
		}
	}

	fmt.Printf("\nTinggi tanaman tertinggi: %.2f cm\n", max)
	fmt.Printf("Tinggi tanaman terpendek: %.2f cm\n", min)
}
