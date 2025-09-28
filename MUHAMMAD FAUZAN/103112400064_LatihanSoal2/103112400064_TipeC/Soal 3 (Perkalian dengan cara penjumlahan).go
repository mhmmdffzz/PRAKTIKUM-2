package main

import (
	"fmt"
)

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)
	hasil := kali(n, m, 0)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
func kali(n, m, hasil int) int {
	if m == 0 {
		return hasil
	}
	return kali(n, m-1, hasil+n)
}
