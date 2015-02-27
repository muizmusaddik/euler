package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	if number == 1 {
		return true
	}

	if number < 0 || number % 2 == 0 {
		return false
	}

	for i := 3; float64(i) <= math.Ceil(math.Sqrt(float64(number))); i += 2 {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func primer(number int) int {
	var largest int
	for i := 2; number > 1; i++ {
		if isPrime(i) {
			if number % i == 0 {
				number = number / i
				largest = i
			}
		} else {
			continue
		}
	}
	return largest
}

func main() {
	fmt.Println(primer(600851475143))
}
