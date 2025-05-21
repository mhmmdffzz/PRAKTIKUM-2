package main

import (
	"fmt"
	"time"
)

func GenerateLastFiveReport(data []Assessment, idUser string) {
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

func GenerateMonthlyAverageReport(data []Assessment, idUser string) float64 {
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

func GenerateRecommendation(skor int) string {
	switch {
	case skor <= 10:
		return "Kondisi stabil. Terus jaga keseimbangan mentalmu!"
	case skor <= 15:
		return "Perlu perhatian. Coba lakukan self-care atau curhat ke teman."
	default:
		return "Skor tinggi. Disarankan berkonsultasi ke profesional."
	}
}
