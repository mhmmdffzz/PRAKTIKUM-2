package main

import (
	"fmt"
)

func main() {
	var dataAssessment []Assessment
	jumlahPertanyaan := 5

	for {
		fmt.Println("\n--- Menu Utama ---")
		fmt.Println("1. Tambah Assessment")
		fmt.Println("2. Ubah Assessment")
		fmt.Println("3. Hapus Assessment")
		fmt.Println("4. Tampilkan 5 Assessment Terakhir")
		fmt.Println("5. Rata-rata Skor 30 Hari Terakhir")
		fmt.Println("6. Urutkan Assessment")
		fmt.Println("7. Cari Assessment")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var idAssessment, idUser, tanggalStr string
			fmt.Print("ID Assessment: ")
			fmt.Scan(&idAssessment)
			fmt.Print("ID Pengguna: ")
			fmt.Scan(&idUser)
			fmt.Print("Tanggal (dd-mm-yyyy): ")
			fmt.Scan(&tanggalStr)
			tanggal := FormatTanggal(tanggalStr)

			jawaban := InputJawabanKuesioner(jumlahPertanyaan)
			skorTotal := HitungTotalSkor(jawaban)

			assessmentBaru := Assessment{
				IDAssessment: idAssessment,
				IDUser:       idUser,
				Tanggal:      tanggal,
				Jawaban:      jawaban,
				SkorTotal:    skorTotal,
			}

			TambahAssessment(&dataAssessment, assessmentBaru)

		case 2:
			var idAssessment string
			fmt.Print("Masukkan ID Assessment yang ingin diubah: ")
			fmt.Scan(&idAssessment)

			var idUser, tanggalStr string
			fmt.Print("ID Pengguna Baru: ")
			fmt.Scan(&idUser)
			fmt.Print("Tanggal Baru (dd-mm-yyyy): ")
			fmt.Scan(&tanggalStr)
			tanggal := FormatTanggal(tanggalStr)

			jawaban := InputJawabanKuesioner(jumlahPertanyaan)
			skorTotal := HitungTotalSkor(jawaban)

			assessmentBaru := Assessment{
				IDAssessment: idAssessment,
				IDUser:       idUser,
				Tanggal:      tanggal,
				Jawaban:      jawaban,
				SkorTotal:    skorTotal,
			}

			UbahAssessment(&dataAssessment, idAssessment, assessmentBaru)

		case 3:
			fmt.Println("progress")

		case 4:
			fmt.Println("progress")

		case 5:
			fmt.Println("progress")

		case 6:
			fmt.Println("progress")

		case 7:
			fmt.Println("progress")

		case 0:
			fmt.Println("Thanks")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
