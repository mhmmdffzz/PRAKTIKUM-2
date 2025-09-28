//MUHAMMAD FAUZAN
//103112400064

package main

import "fmt"

func nilaiFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return nilaiFibonacci(n-1) + nilaiFibonacci(n-2)
}

func main() {
	var a int
	fmt.Scan(&a)

	for indeks := 0; indeks <= 10; indeks++ {
		fmt.Printf("Fibonacci(%d) = %d\n", indeks, nilaiFibonacci(indeks))
	}
}
