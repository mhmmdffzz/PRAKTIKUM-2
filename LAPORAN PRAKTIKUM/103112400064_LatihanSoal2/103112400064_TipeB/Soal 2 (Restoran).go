package main

import "fmt"

func main() {
	var m int
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		var menu, orang, sisa int
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya):\n: ")
		fmt.Scan(&menu, &orang, &sisa)

		total := hitungBiaya(menu, orang, sisa)
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, total)
	}
}

func hitungBiaya(menu, orang, sisa int) int {
	tarif := 0
	if menu <= 3 {
		tarif = 10000
	} else if menu <= 50 {
		tarif = 10000 + (menu-3)*2500
	} else {
		tarif = 100000
	}

	if sisa == 1 {
		tarif *= orang
	}
	return tarif
}
