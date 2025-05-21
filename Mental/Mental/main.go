package main

import (
	"fmt"
	"time"
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
		fmt.Println("8. Tampilkan Rekomendasi Terbaru")
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
			var idAssessment string
			fmt.Print("Masukkan ID Assessment yang ingin dihapus: ")
			fmt.Scan(&idAssessment)
			HapusAssessment(&dataAssessment, idAssessment)

		case 4:
			var idUser string
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			TampilkanLimaTerakhir(dataAssessment, idUser)

		case 5:
			var idUser string
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			rata := HitungRataRataSebulan(dataAssessment, idUser)
			if rata == 0 {
				fmt.Println("Tidak ada data dalam 30 hari terakhir.")
			} else {
				fmt.Printf("Rata-rata skor: %.2f\n", rata)
			}

		case 6:
			fmt.Println("1. Urutkan berdasarkan skor total (Selection Sort)")
			fmt.Println("2. Urutkan berdasarkan tanggal (Insertion Sort)")
			fmt.Print("Pilih metode pengurutan: ")
			var opsi int
			fmt.Scan(&opsi)

			switch opsi {
			case 1:
				SelectionSortBySkor(dataAssessment)
				fmt.Println("Data telah diurutkan berdasarkan skor total.")
			case 2:
				InsertionSortByTanggal(dataAssessment)
				fmt.Println("Data telah diurutkan berdasarkan tanggal.")
			default:
				fmt.Println("Pilihan tidak valid.")
			}

		case 7:
			var idUser string
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			hasil := SequentialSearch(dataAssessment, idUser)
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

		case 8:
			var idUser string
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			hasil := SequentialSearch(dataAssessment, idUser)
			if len(hasil) == 0 {
				fmt.Println("Data tidak ditemukan.")
			} else {
				skor := hasil[len(hasil)-1].SkorTotal
				fmt.Println("Skor terakhir:", skor)
				fmt.Println("Rekomendasi:", Rekomendasi(skor))
			}

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ----------------- FUNGSI TAMBAHAN -----------------

func TampilkanLimaTerakhir(data []Assessment, idUser string) {
	count := 0
	fmt.Println("\n5 Assessment Terakhir:")
	for i := len(data) - 1; i >= 0 && count < 5; i-- {
		if data[i].IDUser == idUser {
			fmt.Printf("%d. Tanggal: %s, Skor: %d\n", count+1,
				data[i].Tanggal.Format("02-01-2006"),
				data[i].SkorTotal)
			count++
		}
	}
	if count == 0 {
		fmt.Println("Data tidak ditemukan.")
	}
}

func HitungRataRataSebulan(data []Assessment, idUser string) float64 {
	now := time.Now()
	var total, count int
	for _, a := range data {
		if a.IDUser == idUser && now.Sub(a.Tanggal).Hours() <= 24*30 {
			total += a.SkorTotal
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count)
}

func SelectionSortBySkor(data []Assessment) {
	n := len(data)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].SkorTotal < data[min].SkorTotal {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func InsertionSortByTanggal(data []Assessment) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Tanggal.After(key.Tanggal) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func SequentialSearch(data []Assessment, idUser string) []Assessment {
	var hasil []Assessment
	for _, a := range data {
		if a.IDUser == idUser {
			hasil = append(hasil, a)
		}
	}
	return hasil
}

func Rekomendasi(skor int) string {
	switch {
	case skor <= 10:
		return "Kondisi stabil. Terus jaga keseimbangan mentalmu!"
	case skor <= 15:
		return "Perlu perhatian. Coba lakukan self-care atau curhat ke teman."
	default:
		return "Skor tinggi. Disarankan berkonsultasi ke profesional."
	}
}
