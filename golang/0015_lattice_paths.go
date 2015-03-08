package main

import "fmt"

func lattice_paths(grid float64) (path float64) {

	//n = Case of square, maximum number of grid from top-left to lowest-right is gridxgrid = 2grid
	//r = Minimal route to go up is always the height of grid
	//Number of possible route combination from top-left to lowest-right is C(n, r) = n! / r! * (n - r)!

	factorial := func(number float64) (product float64) {
		product = 1
		for i := number; i >= 1; i-- {
			product *= i
		}
		return
	}

	path = factorial(2*grid) / (factorial(grid) * factorial(grid))
	return

}

func main() {
	fmt.Println(int(lattice_paths(2)))
	fmt.Println(int(lattice_paths(20)))
}
