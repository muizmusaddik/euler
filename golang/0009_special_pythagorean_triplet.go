package main

import (
	"fmt"
	"math"
)

func greatestCommonDevisor(a int, b int, c int) int {
	var tmp_1 int 
	var tmp_2 int

	for c % a != 0 && b % a != 0 {
		tmp_1 = c
		tmp_2 = b
		c = a
		b = tmp_1 % a
		a = tmp_2 % a
	}

	return a
}

func isCoprime(i int, j int) bool {
	var biggest int

	if i > j {
		biggest = i
	} else {
		biggest = j
	}

	for k := 2; k <= biggest; k++ {
		if i%k == 0 && j%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Euclid's formula
	// a + b + c = m^2 - n^2 + 2mn + (m^2 + n^2) = 1000
	// 2 * m * (m + n) = 1000

	var total int = 1000
	for m := 2; m <= total; m++ {
		for n := 1; n <= m; n++ {
			if 2 * m * (m + n) == total {
				a, b, c := math.Pow(float64(m), 2)- math.Pow(float64(n), 2), 2 * float64(m) * float64(n), math.Pow(float64(m), 2) + math.Pow(float64(n), 2)
				fmt.Println(a, b, c, int(a * b * c))
			}
		}
	}
}

