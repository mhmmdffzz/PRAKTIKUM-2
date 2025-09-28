package main

import "fmt"

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scanln(&bilangan)
	pecahDanJumlah(bilangan)
}

func pecahDanJumlah(bilangan int) {
	var b, angka1, angka2, hasil, pangkat, digit int
	b = bilangan

	for b > 0 {
		digit++
		b /= 10
	}
	ambilDepan := (digit + 1) / 2
	pangkat = 1
	for i := 0; i < digit-ambilDepan; i++ {
		pangkat *= 10
	}

	angka1 = bilangan / pangkat
	angka2 = bilangan % pangkat
	hasil = angka1 + angka2
	fmt.Println("Bilangan 1:", angka1)
	fmt.Println("Bilangan 2:", angka2)
	fmt.Println("Hasil penjumlahan:", hasil)
}
