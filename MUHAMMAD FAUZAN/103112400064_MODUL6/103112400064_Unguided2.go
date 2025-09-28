// MUHAMMAD FAUZAN
// 103112400064
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapusIdx, cari int
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)
	data := make([]int, n)
	fmt.Println("Masukkan elemen:")
	for i := 0; i < n; i++ {
		fmt.Printf("indeks ke-%d: ", i)
		fmt.Scan(&data[i])
	}

	// a. Tampilkan seluruh isi array
	fmt.Print("a. Isi array: ")
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// b. Indeks ganjil
	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// c. Indeks genap
	fmt.Print("c. Indeks genap: ")
	for i := 0; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// d. Indeks kelipatan x
	fmt.Print("Masukkan Indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("d. Indeks kelipatan %d: ", x)
	for i := 0; i < len(data); i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	// e. Hapus elemen pada indeks tertentu
	fmt.Print("Masukan Indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	data = append(data[:hapusIdx], data[hapusIdx+1:]...) // hapus elemen

	fmt.Print("e. Array Setelah dihapus: ")
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// f. Rata-rata
	var total float64
	for _, v := range data {
		total += float64(v)
	}
	rata := total / float64(len(data))
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	// g. Simpangan baku
	var jumlahKuadrat float64
	for _, v := range data {
		selisih := float64(v) - rata
		jumlahKuadrat += selisih * selisih
	}
	sd := math.Sqrt(jumlahKuadrat / float64(len(data)))
	fmt.Printf("g. Simpangan baku: %.2f\n", sd)

	// h. Frekuensi bilangan tertentu
	fmt.Print("Masukan Nilai yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	freq := 0
	for _, v := range data {
		if v == cari {
			freq++
		}
	}
	fmt.Printf("h. Frekuensi %d: %d\n", cari, freq)
}
