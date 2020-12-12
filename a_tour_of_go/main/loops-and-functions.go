/*
Drive into the /main folder

> go run loops-and-functions.go

Reference: https://tour.golang.org/flowcontrol/8
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Possible initial guess {1, x, x/2}
	z := 1.0
	oldZ := 0.0

	/*
		for i:=0; i < 10; i++ {
			z -= (z*z - x) / (2*z)
		} */

	shortFloat := func(number float64) float64 {
		dec := 1000.0
		return math.Round(number*dec) / dec
	}

	for shortFloat(z) != shortFloat(oldZ) {
		oldZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, oldZ)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
