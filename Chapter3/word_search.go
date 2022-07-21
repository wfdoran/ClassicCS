package main

import (
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

func main() {
	g := NewGrid(10, 10)
	fmt.Println(g)
}
