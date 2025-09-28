// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin, *bMax = arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan berat masing-masing balita:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, n, &min, &max)
	rata := rerata(data, n)

	fmt.Printf("Berat minimum: %.2f Kg\n", min)
	fmt.Printf("Berat maksimum: %.2f Kg\n", max)
	fmt.Printf("Rata-rata berat: %.2f Kg\n", rata)
}
