package main

import (
	"fmt"
	"time"
)

// InputJawabanKuesioner meminta input jawaban kuesioner dengan validasi
func InputJawabanKuesioner(jumlahPertanyaan int) []int {
	jawaban := make([]int, jumlahPertanyaan)
	for i := 0; i < jumlahPertanyaan; i++ {
		inputValid := false
		for !inputValid {
			fmt.Printf("Pertanyaan %d: ", i+1)
			var input int
			fmt.Scan(&input)
			if input >= 1 && input <= 5 {
				jawaban[i] = input
				inputValid = true // Keluar dari loop saat input valid
			} else {
				fmt.Println("Error: Masukkan angka 1-5")
			}
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
