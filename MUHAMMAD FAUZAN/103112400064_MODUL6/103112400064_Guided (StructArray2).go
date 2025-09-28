package main

import (
	"fmt"
)

func main() {
	// Deklarasi dan inisialisasi array nilai mahasiswa
	nilaiMahasiswa := [5]int{85, 90, 78, 88, 95}

	fmt.Println("Data Nilai Mahasiswa:")
	fmt.Println("=====================")

	// Menampilkan nilai per mahasiswa
	for i, nilai := range nilaiMahasiswa {
		fmt.Printf("Mahasiswa %d: %d\n", i+1, nilai)
	}

	// Menghitung rata-rata nilai
	var total int
	for _, nilai := range nilaiMahasiswa {
		total += nilai
	}
	rataRata := float64(total) / float64(len(nilaiMahasiswa))

	fmt.Println("=====================")
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata)

	// Mencari nilai tertinggi dan terendah
	tertinggi := nilaiMahasiswa[0]
	terendah := nilaiMahasiswa[0]

	for _, nilai := range nilaiMahasiswa {
		if nilai > tertinggi {
			tertinggi = nilai
		}
		if nilai < terendah {
			terendah = nilai
		}
	}

	fmt.Printf("Nilai tertinggi: %d\n", tertinggi)
	fmt.Printf("Nilai terendah: %d\n", terendah)

	// Contoh array 2 dimensi
	fmt.Println("\nContoh Array 2 Dimensi:")
	fmt.Println("=====================")

	// Nilai ujian mahasiswa dalam 2 mata kuliah (Matematika, Bahasa)
	nilaiUjian := [3][2]int{
		{80, 85},
		{90, 75},
		{70, 95},
	}

	// Menampilkan nilai ujian per mahasiswa
	fmt.Println("Nilai Ujian Mahasiswa (Matematika, Bahasa):")
	for i, nilai := range nilaiUjian {
		fmt.Printf("Mahasiswa %d: Matematika = %d, Bahasa = %d\n", i+1, nilai[0], nilai[1])
	}
}
