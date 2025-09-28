package main

import "fmt"

func pangkatRekursif(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * pangkatRekursif(base, exp-1)
}
func faktorialRekursif(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorialRekursif(n-1)
}
func main() {
	var base, exp, n int
	// Input pangkat
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatRekursif(base, exp))

	// Input faktorial
	fmt.Print("Masukkan angka untuk faktorial: ")
	fmt.Scanln(&n)

	fmt.Printf("%d! = %d\n", n, faktorialRekursif(n))
}
