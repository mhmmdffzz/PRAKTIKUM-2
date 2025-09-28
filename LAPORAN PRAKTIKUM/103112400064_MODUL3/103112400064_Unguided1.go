// MUHAMMAD FAUZAN
//103112400064

package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(permutasi(a, c), kombinasi(a, c))
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("tidak sesuai")
	}
}

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
