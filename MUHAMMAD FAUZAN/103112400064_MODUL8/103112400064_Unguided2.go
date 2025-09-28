// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

func main() {
	const N = 20
	var suara [N + 1]int
	var input, total, valid int
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		total++
		if input >= 1 && input <= N {
			suara[input]++
			valid++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", valid)

	var ketua, wakil int
	for i := 1; i <= N; i++ {
		if suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil && i != ketua) {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
