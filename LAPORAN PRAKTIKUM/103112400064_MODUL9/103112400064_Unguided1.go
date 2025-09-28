// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

func selectionSort(angka []int, panjang int) {
	for i := 0; i < panjang-1; i++ {
		indeksTerkecil := i
		for j := i + 1; j < panjang; j++ {
			if angka[j] < angka[indeksTerkecil] {
				indeksTerkecil = j
			}
		}
		angka[i], angka[indeksTerkecil] = angka[indeksTerkecil], angka[i]
	}
}

func cariMedian(angkaTerurut []int, panjang int) int {
	if panjang%2 == 1 {
		return angkaTerurut[panjang/2]
	}
	return (angkaTerurut[panjang/2-1] + angkaTerurut[panjang/2]) / 2
}

func main() {
	var masukan int
	var daftarAngka [1000000]int
	jumlahAngka := 0
	selesai := false

	for !selesai {
		fmt.Scan(&masukan)

		if masukan == -5313 {
			selesai = true

		} else if masukan == 0 {
			var salinanAngka [1000000]int
			for i := 0; i < jumlahAngka; i++ {
				salinanAngka[i] = daftarAngka[i]
			}

			selectionSort(salinanAngka[:], jumlahAngka)
			fmt.Println(cariMedian(salinanAngka[:], jumlahAngka))

		} else if masukan > 0 {
			daftarAngka[jumlahAngka] = masukan
			jumlahAngka++
		}
	}
}
