// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

func main() {
	var berat [1000]float64
	var n int

	fmt.Print("Masukkan jumlah anak kelinci (maks 1000): ")
	fmt.Scan(&n)

	if n < 1 || n > 1000 {
		fmt.Println("Jumlah tidak valid! Harus antara 1 hingga 1000.")
		return
	}

	fmt.Println("\nMasukkan berat masing-masing anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min, max := berat[0], berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("\nBerat terkecil: %.2f kg\n", min)
	fmt.Printf("Berat terbesar: %.2f kg\n", max)
}
