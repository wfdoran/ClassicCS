package main

import (
	"classic_sc/csp"
	"fmt"
)

type ChipLocation struct {
	row_start int
	col_start int
	row_end   int
	col_end   int
}

func GenerateChipLocations(num_rows int, num_cols int) chan ChipLocation {
	ch := make(chan ChipLocation)

	go func(ch chan ChipLocation) {
		for rs := 0; rs < num_rows; rs++ {
			for re := rs; re < num_rows; re++ {
				for cs := 0; cs < num_cols; cs++ {
					for ce := cs; ce < num_cols; ce++ {
						loc := ChipLocation{rs, cs, re, ce}
						ch <- loc
					}
				}
			}
		}
		close(ch)
	}(ch)

	return ch
}

type Chip struct {
	num_rows int
	num_cols int
}

func gen_size_check(chip Chip) func(assignment map[Chip]ChipLocation) bool {
	return func(assignment map[Chip]ChipLocation) bool {
		loc, ok := assignment[chip]

		if !ok {
			return true
		}

		loc_num_rows := loc.row_end - loc.row_start + 1
		loc_num_cols := loc.col_end - loc.col_start + 1

		if chip.num_rows == loc_num_rows && chip.num_cols == loc_num_cols {
			return true
		}
		if chip.num_cols == loc_num_rows && chip.num_rows == loc_num_cols {
			return true
		}

		return false
	}
}

func gen_overlap_check(chip1 Chip, chip2 Chip) func(assignment map[Chip]ChipLocation) bool {
	return func(assignment map[Chip]ChipLocation) bool {
		loc1, ok1 := assignment[chip1]
		loc2, ok2 := assignment[chip2]

		if !ok1 || !ok2 {
			return true
		}

		temp := make(map[[2]int]bool)

		for r := loc1.row_start; r <= loc1.row_end; r++ {
			for c := loc1.col_start; c <= loc1.col_end; c++ {
				temp[[2]int{r, c}] = true
			}
		}

		for r := loc2.row_start; r <= loc2.row_end; r++ {
			for c := loc2.col_start; c <= loc2.col_end; c++ {
				_, ok := temp[[2]int{r, c}]
				if ok {
					return false
				}
			}
		}

		return true
	}
}

func main() {
	csp := csp.New[Chip, ChipLocation]()

	num_rows := 9
	num_cols := 9

	for d := range GenerateChipLocations(num_rows, num_cols) {
		csp.AddDomain(d)
	}

	chips := []Chip{Chip{1, 6}, Chip{4, 4}, Chip{3, 3}, Chip{2, 2}, Chip{2, 5}}

	for _, chip := range chips {
		csp.AddVariable(chip)
	}

	for _, chip := range chips {
		inputs := []Chip{chip}
		check := gen_size_check(chip)

		csp.AddConstraintGeneral(inputs, check)
	}

	for i := 0; i < len(chips); i++ {
		for j := i + 1; j < len(chips); j++ {
			chip1 := chips[i]
			chip2 := chips[j]

			inputs := []Chip{chip1, chip2}
			check := gen_overlap_check(chip1, chip2)

			csp.AddConstraintGeneral(inputs, check)
		}
	}

	soln := csp.BacktrackSearch()

	// fmt.Println(soln)

	board := make(map[[2]int]int)
	for chip_idx, chip := range chips {
		loc, _ := soln[chip]

		for r := loc.row_start; r <= loc.row_end; r++ {
			for c := loc.col_start; c <= loc.col_end; c++ {
				board[[2]int{r, c}] = chip_idx
			}
		}
	}

	for r := 0; r < num_rows; r++ {
		for c := 0; c < num_cols; c++ {
			chip_idx, ok := board[[2]int{r, c}]
			if ok {
				fmt.Printf("%d", chip_idx)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
