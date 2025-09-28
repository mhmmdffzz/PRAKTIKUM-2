// MUHAMMAD FAUZAN
//103112400064

package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	if jarak(cx1, cy1, x, y) <= r1 && jarak(cx2, cy2, x, y) <= r2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if jarak(cx1, cy1, x, y) <= r1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if jarak(cx2, cy2, x, y) <= r2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
