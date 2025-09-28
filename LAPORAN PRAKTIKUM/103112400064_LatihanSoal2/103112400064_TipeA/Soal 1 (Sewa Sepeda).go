package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var jam, menit int
	var member bool
	var voucher string
	fmt.Print("Jam: ")
	fmt.Scanln(&jam)
	fmt.Print("Menit: ")
	fmt.Scanln(&menit)
	fmt.Print("Member (true/false): ")
	fmt.Scanln(&member)
	fmt.Print("Voucher: ")
	fmt.Scanln(&voucher)

	totalJam := durasi(jam, menit)
	tarif := 5000.0
	if member {
		tarif = 3240.74
	}

	biaya := tarif * float64(totalJam)
	if totalJam >= 3 && diskon(voucher) {
		biaya -= biaya * 0.10
	}
	biaya = math.Round(biaya*100) / 100
	fmt.Printf("Biaya: Rp %.2f\n", biaya)
}

func durasi(j, m int) int {
	if j == 0 && m < 10 {
		return 0
	}
	if m >= 10 {
		return j + 1
	}
	return j
}

func diskon(v string) bool {
	for _, c := range v {
		if c == '5' || c == '6' {
			return true
		}
	}
	return false
}
