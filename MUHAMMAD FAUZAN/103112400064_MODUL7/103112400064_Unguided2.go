// MUHAMMAD FAUZAN
// 103112400064
package main

import (
	"fmt"
)

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah (1 dan 2): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Nilai x harus 1–1000 dan y > 0")
		return
	}

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalWadah := make([]float64, jumlahWadah)

	for idx := 0; idx < jumlahWadah; idx++ {
		start := idx * y
		end := start + y
		if end > x {
			end = x
		}
		for i := start; i < end; i++ {
			totalWadah[idx] += berat[i]
		}
	}

	fmt.Println("Total berat tiap wadah:")
	totalBeratSemua := 0.0
	for _, total := range totalWadah {
		fmt.Printf("%.2f ", total)
		totalBeratSemua += total
	}
	fmt.Println()

	rataRata := totalBeratSemua / float64(jumlahWadah)
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRata)
}
