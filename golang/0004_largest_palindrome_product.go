package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

func isPalindromeNumber(number int) bool {
	arrays := strings.Split(strconv.Itoa(number), "")
	length := math.Floor(float64(len(arrays) / 2))
	for i := 0; i < int(length); i++ {
		if arrays[i] != arrays[len(arrays) - 1 - i] {
			return false
		}
	}
	return true
}
func main() {
	var found bool = false
	for i := 998001; i > 0; i-- {
		if isPalindromeNumber(i) {
			for j := 999; j > 899; j-- {
				if i % j == 0 {
					tmp := i / j
					for k := j; k > 899; k-- {
						if tmp % k == 0 {
							fmt.Println("Largest palindrome number made from the product of 2 3-digit number is: ", i)
							found = true
							break
						}
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
	}
}

