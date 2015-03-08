package main

import (
	"fmt"
	"strconv"
	"strings"
)

func letterizeThenSum(number int) int {
	var dict = map[string]int{
		"0":  0, // meh
		"1":  3, // one
		"2":  3, // two
		"3":  5, // three
		"4":  4, // four
		"5":  4, // five
		"6":  3, // six
		"7":  5, // seven
		"8":  5, // eight
		"9":  4, // nine
		"10": 3, // ten
		"11": 6, // eleven
		"12": 6, // twelve
		"13": 8, // thirteen
		"14": 8, // fourteen
		"15": 7, // fifteen
		"16": 7, // sixteen
		"17": 9, // seventeen
		"18": 8, // eighteen
		"19": 8, // nineteen
		"+2": 6, // twenty
		"+3": 6, // thirty
		"+4": 5, // foutty
		"+5": 5, // fifty
		"+6": 5, // sixty
		"+7": 7, // seventy
		"+8": 6, // eighty
		"+9": 6, // ninety
	}

	array := strings.Split(strconv.Itoa(number), "")

	switch len(array) {
	case 4:
		return 11 // one thousand
	case 3:
		sum := dict[array[0]] + 7 // hundred
		if array[1] != "0" || array[2] != "0" {
			sum += 3 // and
			if array[1] == "0" {
				sum += dict[array[2]]
			} else {
				i, _ := strconv.Atoi(array[1])
				if i > 1 {
					sum += dict["+"+array[1]] + dict[array[2]]
				} else {
					sum += dict[array[1]+array[2]]
				}
			}
		}
		return sum
	case 2:
		i, _ := strconv.Atoi(array[0])
		if i > 1 {
			return dict["+"+array[0]] + dict[array[1]]
		} else {
			return dict[array[0]+array[1]]
		}
	case 1:
		return dict[array[0]]
	}

	return 0

}

func main() {
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += letterizeThenSum(i)
	}
	fmt.Println(sum)
}
