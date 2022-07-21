package main

import (
	"classic_sc/csp"
	"fmt"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func gen_diag_check(row1 int, row2 int) func(assignment map[int]int) bool {
	return func(assignment map[int]int) bool {
		col1, ok1 := assignment[row1]
		col2, ok2 := assignment[row2]

		if !ok1 || !ok2 {
			return true
		}
		return abs(row1-row2) != abs(col1-col2)
	}
}

func main() {
	n := 8

	csp := csp.New[int, int]()

	for i := 1; i <= n; i++ {
		csp.AddVariable(i)
	}

	for i := 1; i <= n; i++ {
		csp.AddDomain(i)
	}

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			csp.AddConstraintNotEqual(i, j)
		}
	}

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			inputs := []int{i, j}
			check := gen_diag_check(i, j)

			csp.AddConstraintGeneral(inputs, check)
		}
	}

	soln := csp.BacktrackSearch()

	for row := 1; row <= n; row++ {
		col, _ := soln[row]

		for i := 1; i < col; i++ {
			fmt.Printf(".")
		}
		fmt.Printf("Q")
		for i := col + 1; i <= n; i++ {
			fmt.Printf(".")
		}
		fmt.Println()
	}

}
