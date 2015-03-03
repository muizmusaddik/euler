package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	if number == 2 {
		return true
	}
	if number < 2 || number % 2 == 0 {
		return false
	}

	for i := 3; float64(i) <= math.Ceil(math.Sqrt(float64(number))); i += 2 {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var sum int
	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}