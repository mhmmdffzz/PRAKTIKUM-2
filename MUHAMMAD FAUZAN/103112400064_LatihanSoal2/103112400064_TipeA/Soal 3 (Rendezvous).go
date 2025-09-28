package main

import "fmt"

func rendezvous(x, y int) int {
	j := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			j++
		}
	}
	return j
}
func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("jumlah pertemuan dalam setahun: ", rendezvous(x, y))
}
