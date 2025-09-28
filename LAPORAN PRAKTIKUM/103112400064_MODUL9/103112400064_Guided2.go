package main

import "fmt"

type Mahasiswa struct {
	Nama  string
	Nilai int
}

func printData(mahasiswas []Mahasiswa) {
	for _, m := range mahasiswas {
		fmt.Printf("%s: %d ", m.Nama, m.Nilai)
	}
	fmt.Println()
}

func insertionSort(mahasiswas []Mahasiswa) {
	fmt.Println("\nProses Insertion Sort:")
	for i := 1; i < len(mahasiswas); i++ {
		key := mahasiswas[i]
		j := i - 1

		for j >= 0 && mahasiswas[j].Nilai > key.Nilai {
			mahasiswas[j+1] = mahasiswas[j]
			j--
		}
		mahasiswas[j+1] = key

		// Tampilkan kondisi array setelah iterasi
		fmt.Printf("Iterasi ke-%d: ", i)
		printData(mahasiswas)
	}
}

func main() {
	data := []Mahasiswa{
		{"Budi", 75},
		{"Ani", 90},
		{"Dedi", 65},
		{"Citra", 85},
		{"Eka", 70},
	}

	fmt.Println("Data sebelum Insertion Sort:")
	printData(data)

	insertionSort(data)

	fmt.Println("\nData setelah Insertion Sort (berdasarkan Nilai):")
	printData(data)
}
