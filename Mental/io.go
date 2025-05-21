package main

import (
	"fmt"
	"time"
)

// InputJawabanKuesioner meminta input jawaban kuesioner dengan validasi
func InputJawabanKuesioner(jumlahPertanyaan int) []int {
	jawaban := make([]int, jumlahPertanyaan)
	for i := 0; i < jumlahPertanyaan; i++ {
		for {
			fmt.Printf("Pertanyaan %d (1-5): ", i+1)
			var input int
			_, err := fmt.Scan(&input)
			if err != nil || input < 1 || input > 5 {
				fmt.Println("Error: Masukkan angka antara 1-5")
				continue
			}
			jawaban[i] = input
			break
		}
	}
	return jawaban
}

// HitungTotalSkor menghitung total skor dari jawaban kuesioner
func HitungTotalSkor(jawaban []int) int {
	total := 0
	for _, nilai := range jawaban {
		total += nilai
	}
	return total
}

// FormatTanggal mengkonversi string ke time.Time dengan format dd-mm-yyyy
func FormatTanggal(tanggalString string) (time.Time, error) {
	tanggal, err := time.Parse("02-01-2006", tanggalString)
	if err != nil {
		return time.Time{}, fmt.Errorf("format tanggal salah, gunakan dd-mm-yyyy")
	}
	return tanggal, nil
}
