package main

import "fmt"

func fibonacci() []int {
	var fib []int
	var sum int = 0
	for i := 0; sum < 4000000; i++ {
		if i < 2 {
			fib = append(fib, 1)
		} else {
			fib = append(fib, (fib[i - 1] + fib[i - 2]))
			sum = fib[i] + fib[i - 1]
		}
	}
	return fib
}
func main() {
	var sum int
	for _,val := range fibonacci() {
		if val % 2 == 0 {
			sum += val
		}
	}

	fmt.Println(sum)
}
