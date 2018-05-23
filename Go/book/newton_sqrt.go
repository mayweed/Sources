package main

import (
	"fmt"
)

//Go Tour 2d part chap 8
// idea:return the value for each iteration of the loop
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / 2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
