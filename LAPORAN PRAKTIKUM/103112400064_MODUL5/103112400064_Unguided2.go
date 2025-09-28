//MUHAMMAD FAUZAN
//103112400064

package main

import "fmt"

func printStars(n int) {
	if n == 0 {
		return
	}
	printStars(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var N int
	fmt.Scan(&N)
	printStars(N)
}
