// MUHAMMAD FAUZAN
// 103112400064
package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	*n = 0
	fmt.Print("Teks : ")
	for {
		fmt.Scanf("%c", &karakter)

		if karakter == '.' || *n >= NMAX {
			break
		}
		if karakter != ' ' && karakter != '\n' && karakter != '\r' {
			t[*n] = karakter
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var salin tabel
	for i := 0; i < n; i++ {
		salin[i] = t[i]
	}
	balikanArray(&salin, n)

	for i := 0; i < n; i++ {
		if t[i] != salin[i] {
			return false
		}
	}
	return true
}

func main() {
	var teks tabel
	var jumlah int
	isiArray(&teks, &jumlah)
	fmt.Print("Teks : ")
	cetakArray(teks, jumlah)

	if palindrom(teks, jumlah) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}
