package main

import (
	"classic_sc/csp"
	"fmt"
)

func check_sum(assignment map[rune]int) bool {
	d, ok_d := assignment['D']
	e, ok_e := assignment['E']
	y, ok_y := assignment['Y']

	if !ok_d || !ok_e || !ok_y {
		return true
	}

	term1 := d
	term2 := e
	term3 := y

	if (term1+term2)%10 != term3 {
		return false
	}

	n, ok_n := assignment['N']
	r, ok_r := assignment['R']

	if !ok_n || !ok_r {
		return true
	}

	term1 += 10 * n
	term2 += 10 * r
	term3 += 10 * e

	if (term1+term2)%100 != term3 {
		return false
	}

	o, ok_o := assignment['O']

	if !ok_o {
		return true
	}

	term1 += 100 * e
	term2 += 100 * o
	term3 += 100 * n

	if (term1+term2)%1000 != term3 {
		return false
	}

	s, ok_s := assignment['S']
	m, ok_m := assignment['M']

	if !ok_s || !ok_m {
		return true
	}

	term1 += 1000 * s
	term2 += 1000 * m
	term3 += 1000*o + 10000*m

	return term1+term2 == term3
}

func gen_not_zero(c rune) func(assignment map[rune]int) bool {
	return func(assignment map[rune]int) bool {
		d, ok := assignment[c]

		if !ok {
			return true
		}

		return d != 0
	}
}

func main() {
	csp := csp.New[rune, int]()

	vars := []rune{'S', 'E', 'N', 'D', 'M', 'O', 'R', 'Y'}
	for _, v := range vars {
		csp.AddVariable(v)
	}

	for i := 0; i <= 9; i++ {
		csp.AddDomain(i)
	}

	num_vars := len(vars)
	for i := 0; i < num_vars; i++ {
		for j := i + 1; j < num_vars; j++ {
			csp.AddConstraintNotEqual(vars[i], vars[j])
		}
	}

	csp.AddConstraintGeneral(vars, check_sum)

	csp.AddConstraintGeneral([]rune{'S'}, gen_not_zero('S'))
	csp.AddConstraintGeneral([]rune{'M'}, gen_not_zero('M'))

	soln := csp.BacktrackSearch()

	d, _ := soln['D']
	e, _ := soln['E']
	y, _ := soln['Y']
	n, _ := soln['N']
	r, _ := soln['R']
	o, _ := soln['O']
	s, _ := soln['S']
	m, _ := soln['M']

	fmt.Println(" ", s, e, n, d)
	fmt.Println(" ", m, o, r, e)
	fmt.Println("-----")
	fmt.Println(m, o, n, e, y)

}
