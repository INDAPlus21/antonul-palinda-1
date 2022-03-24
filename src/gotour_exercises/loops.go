package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	dif := 1.0
	for math.Abs(dif) > 0.0000000001 {
		dif = (z*z - x) / (2 * z)
		z -= dif
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
