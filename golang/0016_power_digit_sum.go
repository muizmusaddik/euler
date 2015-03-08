package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sum(power float64) int {
	var sum int
	array := strings.Split(strconv.FormatFloat(math.Pow(2, power), 'f', 0, 64), "")

	for _, i := range array {
		v, _ := strconv.Atoi(i)
		sum += v
	}

	return sum

}

func main() {
	fmt.Println(sum(1000))
}
