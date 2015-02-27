package main

import "fmt"

func sum(number int) int {
	var total int
	for	 i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func main() {
	fmt.Println (sum(1000))
}
