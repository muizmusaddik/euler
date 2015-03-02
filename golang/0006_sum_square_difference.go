package main

import (
	"fmt"
	"math"
	"strconv"
)

func sumOfSquaresAndSquareOfSums(number int) (float64, float64) {
	var sum int
	var squares float64
	for i := 1; i <= number; i++ {
		sum += i
		squares += math.Pow(float64(i), 2)
	}

	return math.Pow(float64(sum), 2), squares
}

func main() {
	squares, sum := sumOfSquaresAndSquareOfSums(100)
	fmt.Println(strconv.FormatFloat((squares - sum), 'f', 0, 64))
}