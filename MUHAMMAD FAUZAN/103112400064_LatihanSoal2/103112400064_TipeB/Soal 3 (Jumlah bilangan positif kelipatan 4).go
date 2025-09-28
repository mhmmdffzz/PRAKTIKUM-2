package main

import (
	"fmt"
)

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	fmt.Print("Masukkan bilangan (negatif untuk berhenti):\n: ")
	total := hasil(0)
	fmt.Println("Jumlah bilangan kelipatan 4:", total)
}

func hasil(total int) int {
	var a int
	fmt.Scan(&a)
	if a < 0 {
		return total
	}
	if a > 0 && a%4 == 0 {
		total += a
	}
	return hasil(total)
}
