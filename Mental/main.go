package main

import (
	"fmt"
)

func main() {
	var dataAssessment []Assessment
	jumlahPertanyaan := 5

	for {
		tampilkanMenuUtama()

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAssessmentBaru(&dataAssessment, jumlahPertanyaan)
		case 2:
			ubahAssessment(&dataAssessment, jumlahPertanyaan)
		case 3:
			hapusAssessment(&dataAssessment)
		case 4:
			tampilkanLaporanLimaTerakhir(dataAssessment)
		case 5:
			tampilkanRataRataSebulan(dataAssessment)
		case 6:
			urutkanAssessment(&dataAssessment)
		case 7:
			cariAssessment(dataAssessment)
		case 8:
			tampilkanRekomendasiTerbaru(dataAssessment)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanMenuUtama() {
	fmt.Println("\n--- Menu Utama ---")
	fmt.Println("1. Tambah Assessment")
	fmt.Println("2. Ubah Assessment")
	fmt.Println("3. Hapus Assessment")
	fmt.Println("4. Tampilkan 5 Assessment Terakhir")
	fmt.Println("5. Rata-rata Skor 30 Hari Terakhir")
	fmt.Println("6. Urutkan Assessment")
	fmt.Println("7. Cari Assessment")
	fmt.Println("8. Tampilkan Rekomendasi Terbaru")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahAssessmentBaru(dataAssessment *[]Assessment, jumlahPertanyaan int) {
	var idAssessment, idUser, tanggalStr string
	fmt.Print("ID Assessment: ")
	fmt.Scan(&idAssessment)
	fmt.Print("ID Pengguna: ")
	fmt.Scan(&idUser)
	fmt.Print("Tanggal (dd-mm-yyyy): ")
	fmt.Scan(&tanggalStr)

	tanggal, err := FormatTanggal(tanggalStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jawaban := InputJawabanKuesioner(jumlahPertanyaan)
	skorTotal := HitungTotalSkor(jawaban)

	assessmentBaru := Assessment{
		IDAssessment: idAssessment,
		IDUser:       idUser,
		Tanggal:      tanggal,
		Jawaban:      jawaban,
		SkorTotal:    skorTotal,
	}

	TambahAssessment(dataAssessment, assessmentBaru)
}

func ubahAssessment(dataAssessment *[]Assessment, jumlahPertanyaan int) {
	var idAssessment string
	fmt.Print("Masukkan ID Assessment yang ingin diubah: ")
	fmt.Scan(&idAssessment)

	var idUser, tanggalStr string
	fmt.Print("ID Pengguna Baru: ")
	fmt.Scan(&idUser)
	fmt.Print("Tanggal Baru (dd-mm-yyyy): ")
	fmt.Scan(&tanggalStr)

	tanggal, err := FormatTanggal(tanggalStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jawaban := InputJawabanKuesioner(jumlahPertanyaan)
	skorTotal := HitungTotalSkor(jawaban)

	assessmentBaru := Assessment{
		IDAssessment: idAssessment,
		IDUser:       idUser,
		Tanggal:      tanggal,
		Jawaban:      jawaban,
		SkorTotal:    skorTotal,
	}

	UbahAssessment(dataAssessment, idAssessment, assessmentBaru)
}

func hapusAssessment(dataAssessment *[]Assessment) {
	var idAssessment string
	fmt.Print("Masukkan ID Assessment yang ingin dihapus: ")
	fmt.Scan(&idAssessment)
	HapusAssessment(dataAssessment, idAssessment)
}

func tampilkanLaporanLimaTerakhir(data []Assessment) {
	var idUser string
	fmt.Print("Masukkan ID Pengguna: ")
	fmt.Scan(&idUser)
	GenerateLastFiveReport(data, idUser)
}

func tampilkanRataRataSebulan(data []Assessment) {
	var idUser string
	fmt.Print("Masukkan ID Pengguna: ")
	fmt.Scan(&idUser)
	rata := GenerateMonthlyAverageReport(data, idUser)
	if rata == 0 {
		fmt.Println("Tidak ada data dalam 30 hari terakhir.")
	} else {
		fmt.Printf("Rata-rata skor: %.2f\n", rata)
	}
}

func urutkanAssessment(data *[]Assessment) {
	fmt.Println("1. Urutkan berdasarkan skor total (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan tanggal (Insertion Sort)")
	fmt.Print("Pilih metode pengurutan: ")
	var opsi int
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		SelectionSortBySkor(*data)
		fmt.Println("Data telah diurutkan berdasarkan skor total.")
	case 2:
		InsertionSortByTanggal(*data)
		fmt.Println("Data telah diurutkan berdasarkan tanggal.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariAssessment(data []Assessment) {
	fmt.Println("1. Sequential Search (by User ID)")
	fmt.Println("2. Binary Search (by Assessment ID)")
	fmt.Print("Pilih metode pencarian: ")
	var searchOpt int
	fmt.Scan(&searchOpt)

	if searchOpt == 1 {
		var idUser string
		fmt.Print("Masukkan ID Pengguna: ")
		fmt.Scan(&idUser)
		hasil := SequentialSearchByUser(data, idUser)
		if len(hasil) == 0 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Hasil pencarian:")
			for _, a := range hasil {
				fmt.Printf("ID: %s | Tanggal: %s | Skor: %d\n",
					a.IDAssessment,
					a.Tanggal.Format("02-01-2006"),
					a.SkorTotal)
			}
		}
	} else if searchOpt == 2 {
		var idAssessment string
		fmt.Print("Masukkan ID Assessment: ")
		fmt.Scan(&idAssessment)
		hasil, _ := BinarySearchByID(data, idAssessment)
		if hasil.IDAssessment == "" {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Printf("ID: %s | Tanggal: %s | Skor: %d\n",
				hasil.IDAssessment,
				hasil.Tanggal.Format("02-01-2006"),
				hasil.SkorTotal)
		}
	}
}

func tampilkanRekomendasiTerbaru(data []Assessment) {
	var idUser string
	fmt.Print("Masukkan ID Pengguna: ")
	fmt.Scan(&idUser)
	hasil := SequentialSearchByUser(data, idUser)
	if len(hasil) == 0 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		skor := hasil[len(hasil)-1].SkorTotal
		fmt.Println("Skor terakhir:", skor)
		fmt.Println("Rekomendasi:", GenerateRecommendation(skor))
	}
}
