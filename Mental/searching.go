package main

func BinarySearchByID(data []Assessment, idAssessment string) (Assessment, int) {
	// Pastikan data sudah terurut
	InsertionSortByTanggal(data) // atau sorting lain yang sesuai

	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid].IDAssessment == idAssessment {
			return data[mid], mid
		} else if data[mid].IDAssessment < idAssessment {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return Assessment{}, -1
}

func SequentialSearchByUser(data []Assessment, idUser string) []Assessment {
	var hasil []Assessment
	for _, a := range data {
		if a.IDUser == idUser {
			hasil = append(hasil, a)
		}
	}
	return hasil
}
