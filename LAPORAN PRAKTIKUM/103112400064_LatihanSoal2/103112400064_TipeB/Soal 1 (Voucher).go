package main

import "fmt"

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	jumlah := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	fmt.Println("Banyaknya angka ganjil:", jumlah)
}
