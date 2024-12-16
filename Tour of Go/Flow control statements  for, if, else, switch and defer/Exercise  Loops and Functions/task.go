package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var delta = 1e-6
	z := x
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2), Sqrt(2)-math.Sqrt(2))
}
