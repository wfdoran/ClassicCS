package main

import (
	"classic_sc/csp"
	"fmt"
	"math/rand"
)

type Grid struct {
	num_rows int
	num_cols int
	grid     [][]rune
}

func NewGrid(num_rows int, num_cols int) *Grid {
	g := Grid{num_rows: num_rows, num_cols: num_cols, grid: nil}
	g.grid = make([][]rune, num_rows)
	for r := 0; r < num_rows; r++ {
		g.grid[r] = make([]rune, num_cols)
		for c := 0; c < num_cols; c++ {
			g.grid[r][c] = rune(int('A') + rand.Intn(26))
		}
	}

	return &g
}

func (g Grid) String() string {
	rv := ""

	for r := 0; r < g.num_rows; r++ {
		for c := 0; c < g.num_cols; c++ {
			if c != 0 {
				rv += " "
			}
			rv += string(g.grid[r][c])
		}
		rv += "\n"
	}

	return rv
}

type WordLocation struct {
	row_start int
	col_start int
	row_delta int
	col_delta int
	length    int
}

func GenerateWordLocations(num_rows int, num_cols int) chan WordLocation {
	ch := make(chan WordLocation)

	go func(ch chan WordLocation) {
		for r := 0; r < num_rows; r++ {
			for c := 0; c < num_cols; c++ {
				// down
				for i := 0; ; i++ {
					r2 := r + i
					// c2 := c
					if r2 >= num_rows {
						break
					}

					loc := WordLocation{row_start: r, col_start: c, row_delta: 1, col_delta: 0, length: i + 1}
					ch <- loc
				}

				// left
				for i := 0; ; i++ {
					// r2 := r
					c2 := c + i
					if c2 >= num_cols {
						break
					}

					loc := WordLocation{row_start: r, col_start: c, row_delta: 0, col_delta: 1, length: i + 1}
					ch <- loc
				}

				// diagonal down-right
				for i := 0; ; i++ {
					r2 := r + i
					c2 := c + i
					if r2 >= num_rows || c2 >= num_cols {
						break
					}

					loc := WordLocation{row_start: r, col_start: c, row_delta: 1, col_delta: 1, length: i + 1}
					ch <- loc
				}

				// diagonal down-left
				for i := 0; ; i++ {
					r2 := r + i
					c2 := c - i
					if r2 >= num_rows || c2 < 0 {
						break
					}

					loc := WordLocation{row_start: r, col_start: c, row_delta: 1, col_delta: -1, length: i + 1}
					ch <- loc
				}
			}
		}
		close(ch)
	}(ch)

	return ch
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func gen_len_check(w string) func(assignment map[string]WordLocation) bool {
	return func(assignment map[string]WordLocation) bool {
		loc, ok := assignment[w]
		if !ok {
			return true
		}

		return loc.length == len(w)
	}
}

func gen_overlap_check(w1 string, w2 string) func(assignment map[string]WordLocation) bool {
	return func(assignment map[string]WordLocation) bool {
		loc1, ok1 := assignment[w1]
		loc2, ok2 := assignment[w2]
		if !ok1 || !ok2 {
			return true
		}

		temp := make(map[[2]int]bool)

		for i := 0; i < loc1.length; i++ {
			r := loc1.row_start + i*loc1.row_delta
			c := loc1.col_start + i*loc1.col_delta
			temp[[2]int{r, c}] = true
		}

		for i := 0; i < loc2.length; i++ {
			r := loc2.row_start + i*loc2.row_delta
			c := loc2.col_start + i*loc2.col_delta
			_, ok := temp[[2]int{r, c}]
			if ok {
				return false
			}
		}
		return true
	}
}

func main() {
	csp := csp.New[string, WordLocation]()

	num_rows := 10
	num_cols := 10

	for d := range GenerateWordLocations(num_rows, num_cols) {
		csp.AddDomain(d)
	}

	words := []string{"MATTHEW", "JOE", "MARY", "SARAH", "SALLY"}
	for _, w := range words {
		csp.AddVariable(w)
	}

	for _, w := range words {
		inputs := []string{w}
		check := gen_len_check(w)

		csp.AddConstraintGeneral(inputs, check)
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			w1 := words[i]
			w2 := words[j]

			inputs := []string{w1, w2}
			check := gen_overlap_check(w1, w2)

			csp.AddConstraintGeneral(inputs, check)
		}
	}

	soln := csp.BacktrackSearch()

	if soln != nil {
		fmt.Println(soln)
	} else {
		fmt.Println("err")
	}

	g := NewGrid(10, 10)

	for _, w := range words {
		loc, _ := soln[w]

		for i := 0; i < loc.length; i++ {
			r := loc.row_start + i*loc.row_delta
			c := loc.col_start + i*loc.col_delta
			g.grid[r][c] = rune(w[i])
		}
	}

	fmt.Println(g)
}
