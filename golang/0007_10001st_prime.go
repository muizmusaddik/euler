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

func primer(max int) int {
	var counter int
	var number int = 1
	for {

		if isPrime(number) {
			counter++
		} 

		if counter == max {
			return number
		}

		number ++
	}

	return 0
}

func main() {
	fmt.Println(primer(10001))
}