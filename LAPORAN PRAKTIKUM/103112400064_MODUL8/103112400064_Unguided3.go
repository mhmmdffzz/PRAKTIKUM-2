// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)

	pos := posisi(n, k)
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1

	for kr <= kn {
		mid := (kr + kn) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			kr = mid + 1
		} else {
			kn = mid - 1
		}
	}
	return -1
}
