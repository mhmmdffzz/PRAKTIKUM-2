package main

import (
	"fmt"
)

type Mahasiswa struct {
	Nama  string
	Nilai int
}

func printData(mahasiswas []Mahasiswa) {
	for _, m := range mahasiswas {
		fmt.Printf("%s: %d", m.Nama, m.Nilai)
	}
	fmt.Println()
}

func selectionSort(mahasiswas []Mahasiswa) {
	n := len(mahasiswas)
	fmt.Println("\nProses Selection Sort:")
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if mahasiswas[j].Nilai < mahasiswas[minIdx].Nilai {
				minIdx = j
			}
		}

		mahasiswas[i], mahasiswas[minIdx] = mahasiswas[minIdx], mahasiswas[i]
		fmt.Printf("Iterasi ke-%d: ", i+1)
		printData(mahasiswas)

	}
}
func main() {
	data := []Mahasiswa{
		{" Budi", 75},
		{" Ani", 90},
		{" Dedi", 65},
		{" Citra", 85},
		{" Eka", 70},
	}
	fmt.Println("Data Sebelum Selection Sort:")
	printData(data)
	selectionSort(data)
	fmt.Println("\ndata setelah Selection Sort (Berdasarkan Nilai):")
	printData(data)

}
