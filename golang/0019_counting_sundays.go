package main

import (
	"fmt"
	"time"
	)

func main() {
	var counter int
	for i := 1901; i <= 2000; i++ {
		for j := 1; j <= 12; j++ {
			m := time.Month(j)
			d := time.Date(i, m, 1, 0, 0, 0, 0, time.UTC)
			if d.Weekday() == 0 {
				counter ++
			}
		}
	}
	fmt.Println(counter)
}
