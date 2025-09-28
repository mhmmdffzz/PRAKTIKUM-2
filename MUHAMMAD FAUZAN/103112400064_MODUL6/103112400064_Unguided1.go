// MUHAMMAD FAUZAN
// 103112400064
package main

import (
	"fmt"
	"math"
)

func dalamLingkaran(x, y, r, xt, yt int) bool {
	d := math.Sqrt(float64((xt-x)*(xt-x) + (yt-y)*(yt-y)))
	return d <= float64(r)
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var xt, yt int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&xt, &yt)

	dalam1 := dalamLingkaran(x1, y1, r1, xt, yt)
	dalam2 := dalamLingkaran(x2, y2, r2, xt, yt)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
