package main

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

func SelectionSortByUserID(data []Assessment) {
	n := len(data)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].IDUser < data[min].IDUser {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
