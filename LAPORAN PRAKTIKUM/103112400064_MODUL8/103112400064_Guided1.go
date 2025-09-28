package main

import (
	"fmt"
	"sort"
)

func sequentialSearch(arr []float64, target float64) (int, int) {
	iterations := 0
	for i, val := range arr {
		iterations++
		fmt.Printf("sequential step %d: cek arr [%d]=%.f\n", iterations, i, val)
		if val == target {
			return i, iterations
		}
	}
	return -1, iterations
}
func binarySEarch(arr []float64, target float64) (int, int) {
	iterations := 0
	low := 0
	high := len(arr) - 1
	for low <= high {
		iterations++
		mid := (low + high) / 2
		fmt.Printf("binary stop %d: cek arr[%d]= %.f\n", iterations, mid, arr[mid])

		if arr[mid] == target {
			return mid, iterations
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid - 1
		}
	}
	return -1, iterations
}
func main() {
	//array awal
	data := []float64{2, 7, 9, 1, 5, 6, 18, 13, 25, 20}
	target := 13.0

	fmt.Println("sequenial search (data tidak perlu urut):")
	idxseq, iterseq := sequentialSearch(data, target)
	if idxseq != -1 {
		fmt.Printf("hasil : ditemukan di indeks %d dalam %d langkah\n\n", idxseq, iterseq)
	} else {
		fmt.Printf("hasil : tidak ditemukan setelah %d langkah\n", iterseq)
	}

	//Binary earch array diurutkan
	sort.Float64s(data)
	fmt.Println("Binary search (setelah data diurutkan):")
	fmt.Println("Data terurut:", data)

	idxBin, iterBin := binarySEarch(data, target)
	if idxBin != -1 {
		fmt.Printf("Hasil : Ditemukan indeks %d dalam %d langkah\n", idxBin, iterBin)
	} else {
		fmt.Printf("Hasil : Tidak ditemukan Setelah %d langkah\n", iterBin)
	}
}
