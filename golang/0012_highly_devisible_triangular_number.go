package main

import (
	"fmt"
	"math"
)

func countOfFactor(n int) int {
	var counter int
	for i := float64(1); i <= math.Floor(math.Sqrt(float64(n))); i++ {
		if n%int(i) == 0 {
			counter++
			if n/int(i) != int(i) {
				counter++
			}
		}
	}
	return counter
}

func main() {
	var i int
	var sum int
	for {
		i++
		sum += i

		if countOfFactor(sum) >= 500 {
			fmt.Println(sum)
			break
		}
	}
}
