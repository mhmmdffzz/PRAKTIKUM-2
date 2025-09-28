//MUHAMMAD FAUZAN
//103112400064

package main

import "fmt"

func hitungSkor(waktu int, totalWaktu, totalSoal *int) {
	if waktu <= 300 {
		*totalSoal++
		*totalWaktu += waktu
	}
}

func main() {
	var namaPemenang, peserta1, peserta2 string
	var waktu1, waktu2, totalWaktu1, totalWaktu2, totalSoal1, totalSoal2, maxSoal, maxWaktu int
	fmt.Scan(&peserta1)
	if peserta1 == "Selesai" {
		return
	}
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu1)
		hitungSkor(waktu1, &totalWaktu1, &totalSoal1)
	}
	for {
		fmt.Scan(&peserta2)
		if peserta2 == "Selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu2)
			hitungSkor(waktu2, &totalWaktu2, &totalSoal2)
		}
		switch {
		case totalSoal1 < totalSoal2:
			maxSoal, maxWaktu, namaPemenang = totalSoal2, totalWaktu2, peserta2
		case totalSoal1 > totalSoal2:
			maxSoal, maxWaktu, namaPemenang = totalSoal1, totalWaktu1, peserta1
		case totalWaktu1 > totalWaktu2:
			maxSoal, maxWaktu, namaPemenang = totalSoal1, totalWaktu1, peserta1
		default:
			maxSoal, maxWaktu, namaPemenang = totalSoal2, totalWaktu2, peserta2
		}
	}
	fmt.Println(namaPemenang, maxSoal, maxWaktu)
}
