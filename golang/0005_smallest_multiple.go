package main

import "fmt"

func findSmallestMultiple(biggest int) int {
	var number int = 1
	var multiplier int = 1
	for {
		if number % multiplier == 0 {
			multiplier++
		} else {
			number ++
			multiplier = 1
		}

		if multiplier == biggest {
			return number
		}
	}	
}
func main() {
	fmt.Println(findSmallestMultiple(20))
}