package main

import "fmt"

func collatzSequenceCount(n int) int {
	var counter int = 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		counter++
	}

	return counter
}

func main() {
	var maxCount int
	var maxNum int
	for i := 13; i < 1000000; i++ {
		if maxCount < collatzSequenceCount(i) {
			maxCount = collatzSequenceCount(i)
			maxNum = i

		}
	}

	fmt.Println(maxNum, maxCount)

}
