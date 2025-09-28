package main

import (
	"fmt"
	"sort"
)

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

// Sequential Search berdasarkan nama
func SeqSearch_3(T arrMhs, n int, X string) (int, int) {
	var found int = -1
	var j int = 0
	var iterasi int = 0

	for j < n && found == -1 {
		iterasi++
		if T[j].nama == X {
			found = j
		}
		j++
	}
	return found, iterasi
}

// Binary Search berdasarkan NIM (data harus sudah terurut berdasarkan nim)
func BinarySearch_3(T arrMhs, n int, X string) (int, int) {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1
	var iterasi int = 0

	for kr <= kn && found == -1 {
		iterasi++
		med = (kr + kn) / 2
		if X < T[med].nim {
			kn = med - 1
		} else if X > T[med].nim {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found, iterasi
}

func main() {
	var data arrMhs
	n := 10

	// Mengisi data secara manual
	data = arrMhs{
		{nama: "Ari", nim: "2201", kelas: "A", jurusan: "Informatika", ipk: 3.4},
		{nama: "Budi", nim: "2203", kelas: "A", jurusan: "Informatika", ipk: 3.6},
		{nama: "Cici", nim: "2202", kelas: "B", jurusan: "Sistem Informasi", ipk: 3.5},
		{nama: "Dina", nim: "2205", kelas: "A", jurusan: "Informatika", ipk: 3.3},
		{nama: "Eko", nim: "2204", kelas: "B", jurusan: "Sistem Informasi", ipk: 3.7},
		{nama: "Fajar", nim: "2206", kelas: "C", jurusan: "Informatika", ipk: 3.1},
		{nama: "Gita", nim: "2209", kelas: "C", jurusan: "Informatika", ipk: 3.8},
		{nama: "Hana", nim: "2208", kelas: "B", jurusan: "Sistem Informasi", ipk: 3.2},
		{nama: "Iwan", nim: "2207", kelas: "C", jurusan: "Informatika", ipk: 3.0},
		{nama: "Joko", nim: "2210", kelas: "A", jurusan: "Informatika", ipk: 3.9},
	}

	// Pencarian Sequential Search berdasarkan nama
	namaDicari := "Fajar"
	idxSeq, iterSeq := SeqSearch_3(data, n, namaDicari)

	fmt.Printf("Sequential Search - Cari nama '%s'\n", namaDicari)
	if idxSeq != -1 {
		fmt.Printf("Ditemukan di indeks: %d, Iterasi: %d\n", idxSeq, iterSeq)
	} else {
		fmt.Printf("Tidak ditemukan, Iterasi: %d\n", iterSeq)
	}

	// Urutkan data berdasarkan NIM untuk binary search
	sort.Slice(data[:n], func(i, j int) bool {
		return data[i].nim < data[j].nim
	})

	// Pencarian Binary Search berdasarkan NIM
	nimDicari := "2206"
	idxBin, iterBin := BinarySearch_3(data, n, nimDicari)

	fmt.Printf("\nBinary Search - Cari NIM '%s'\n", nimDicari)
	if idxBin != -1 {
		fmt.Printf("Ditemukan di indeks: %d, Iterasi: %d\n", idxBin, iterBin)
	} else {
		fmt.Printf("Tidak ditemukan, Iterasi: %d\n", iterBin)
	}
}
