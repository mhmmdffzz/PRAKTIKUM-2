//MUHAMMAD FAUZAN
//103112400064

package main

import "fmt"

func insertionSort(angka []int) {
	panjang := len(angka)
	for i := 1; i < panjang; i++ {
		kunci := angka[i]
		j := i - 1

		for j >= 0 && angka[j] > kunci {
			angka[j+1] = angka[j]
			j--
		}
		angka[j+1] = kunci
	}
}

func cekJarakTetap(angka []int) (bool, int) {
	if len(angka) < 2 {
		return true, 0
	}

	jarak := angka[1] - angka[0]
	for i := 2; i < len(angka); i++ {
		if angka[i]-angka[i-1] != jarak {
			return false, 0
		}
	}
	return true, jarak
}

func main() {
	var masukan int
	var daftarAngka []int

	for {
		_, err := fmt.Scan(&masukan)
		if err != nil || masukan < 0 {
			break
		}
		daftarAngka = append(daftarAngka, masukan)
	}

	insertionSort(daftarAngka)

	for _, nilai := range daftarAngka {
		fmt.Print(nilai, " ")
	}
	fmt.Println()

	valid, jarak := cekJarakTetap(daftarAngka)
	if valid {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
