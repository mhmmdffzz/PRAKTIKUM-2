// MUHAMMAD FAUZAN
// 103112400064
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		pertandingan++
	}

	for i, v := range hasil {
		if v == "Draw" {
			fmt.Printf("Hasil %d : Draw\n", i+1)
		} else {
			fmt.Printf("Hasil %d : %s\n", i+1, v)
		}
	}
	fmt.Println("pertandingan selesai.")
}
