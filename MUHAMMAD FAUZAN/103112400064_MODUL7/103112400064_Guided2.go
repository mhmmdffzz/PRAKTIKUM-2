package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah buku dan jumlah buku per rak: ")
	fmt.Scan(&x, &y)

	var hargaBuku [500]float64
	fmt.Println("\nMasukkan harga setiap buku (dalam ribuan Rp): ")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargaBuku[i])
	}

	var hargaRataRata []float64
	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]
		}
		hargaRataRata = append(hargaRataRata, total/float64(y))
	}

	min, max := hargaBuku[0], hargaBuku[0]
	for _, harga := range hargaBuku[:x] {
		if harga < min {
			min = harga
		}
		if harga > max {
			max = harga
		}
	}

	fmt.Printf("\nRata-rata harga per rak: ")
	for _, avg := range hargaRataRata {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Printf("\nHarga termahal: %.2f Rp\n", max)
	fmt.Printf("Harga termurah: %.2f Rp\n", min)
}
