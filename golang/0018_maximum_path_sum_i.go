package main

import "fmt"

func maximumPathSum(r1 []int, r2 []int, r3 []int, pos int) int {
	var j, k, l, m, max, index, start, end int

	maxValue := func(r []int) (max int) {
		for _, i := range r {
			if max < i { max = i }
		}
		return 
	}

	if pos < 0 {
		start = 0
		end = len(r1) - 1
	} else {
		start = pos
		end = pos + 1
	}

	for i := start; i <= end; i++ {
		v := r1[i]

		if len(r2) != 0 {
			j, k = r2[i], r2[i + 1]
		} else {
			j, k = 0, 0
		}

		if len(r3) != 0 {
			l, m = r3[i], r3[i + 1]
		} else {
			l, m = 0, 0
		}
		sum := maxValue([]int{v + j + l, v + j + m, v + k + l, v + k + m})
		if max < sum {
			max = sum
			index = i
		}
	}
	return index
}

func main() {
	var arr = [][]int {
		[]int {75},
		[]int {95, 64},
		[]int {17, 47, 82},
		[]int {18, 35, 87, 10},
		[]int {20, 4, 82, 47, 65},
		[]int {19, 1, 23, 75, 3, 34},
		[]int {88, 2, 77, 73, 7, 63, 67},
		[]int {99, 65, 4, 28, 6, 16, 70, 92},
		[]int {41, 41, 26, 56, 83, 40, 80, 70, 33},
		[]int {41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		[]int {53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		[]int {70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		[]int {91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		[]int {63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		[]int {4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}
	var tmp []int
	var max int
	var pos int = -1
	for i := 0; i < len(arr); i++ {
		if len(arr) > i + 2 {
			pos = maximumPathSum(arr[i], arr[i + 1], arr[i + 2], pos)
		} else if len(arr) > i + 1 {
			pos = maximumPathSum(arr[i], arr[i + 1], []int{}, pos)
		} else {
			pos = maximumPathSum(arr[i], []int{}, []int{}, pos)
		}

		tmp = append(tmp, arr[i][pos])
		max += arr[i][pos]
	}

	fmt.Println(max, tmp)
}