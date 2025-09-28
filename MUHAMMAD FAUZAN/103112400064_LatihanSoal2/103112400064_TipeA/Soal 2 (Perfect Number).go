package main

import "fmt"

func main() {
	fmt.Println("Nama: Muhammad Fauzan\nNIM : 103112400064")
	var a, b int
	fmt.Print("Nilai awal: ")
	fmt.Scan(&a)
	fmt.Print("Nilai akhir: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect number dari %d sampai %d:", a, b)
	for i := a; i <= b; i++ {
		s := 0
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				s += j
			}
		}
		if s == i {
			fmt.Printf(" %d", i)
		}
	}
}
