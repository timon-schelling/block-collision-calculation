package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5.0
	fmt.Println("≈ ", calc(math.Pow(100, n), 1))
	fmt.Println("= ", math.Pi*math.Pow(10, n))
}

func calc(m1 float64, m2 float64) int {

	r := math.Sqrt(m1)

	a := r
	b := r

	tanA := math.Sqrt(m1 / m2)
	A := math.Atan(tanA)

	sinB := (math.Sin(A) / a) * b
	B := math.Asin(sinB)

	C := math.Pi - (A + B)

	x := 0

	tau := math.Pi * 2
	for i := C; i < tau; i += C {
		x++
	}

	return x
}
