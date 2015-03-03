package main

import (
	"fmt"
	"math"
)

func greatestCommonDeviser(i float64, j, float64, k float64) float64 {

}

func isCoprime(i float64, j float64) bool {
	var biggest int

	if i > j {
		biggest = int(i)
	} else {
		biggest = int(j)
	}

	for k := 2; k <= int(biggest); k++ {
		if int(i)%k == 0 && int(j)%k == 0 {
			return false
		}
	}
	return true
}

func main() {
	for m := float64(2); m <= float64(100); m++ {
		for n := float64(1); n < m; n++ {
			if isCoprime(n, m) {
				a, b, c := math.Pow(m, 2)-math.Pow(n, 2), 2*m*n, math.Pow(m, 2)+math.Pow(n, 2)
				fmt.Println(a, b, c, ":", a+b+c)
			}
		}
	}
}
