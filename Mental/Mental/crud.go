package main

import "fmt"

func TambahAssessment(dataAssessment *[]Assessment, assessmentBaru Assessment) {
	*dataAssessment = append(*dataAssessment, assessmentBaru)
	fmt.Println("Assessment berhasil ditambahkan!")
}

func UbahAssessment(dataAssessment *[]Assessment, idAssessment string, assessmentBaru Assessment) {
	for i := range *dataAssessment {
		if (*dataAssessment)[i].IDAssessment == idAssessment {
			(*dataAssessment)[i] = assessmentBaru
			fmt.Println("Assessment berhasil diubah!")
			return
		}
	}
	fmt.Println("Assessment tidak ditemukan.")
}

func HapusAssessment(dataAssessment *[]Assessment, idAssessment string) {
	for i := range *dataAssessment {
		if (*dataAssessment)[i].IDAssessment == idAssessment {
			*dataAssessment = append((*dataAssessment)[:i], (*dataAssessment)[i+1:]...)
			fmt.Println("Assessment berhasil dihapus!")
			return
		}
	}
	fmt.Println("Assessment tidak ditemukan.")
}
