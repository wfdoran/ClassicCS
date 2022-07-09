package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell int

const (
	Empty Cell = iota
	Blocked
	Start
	Goal
	Path
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return " "
	case Blocked:
		return "X"
	case Start:
		return "S"
	case Goal:
		return "G"
	case Path:
		return "*"
	default:
		return "err"
	}
}

type MazeLocation struct {
	row int
	col int
}

type Maze struct {
	num_rows int
	num_cols int
	grid     [][]Cell
}

func (m Maze) String() string {
	rv := ""
	rv += "+"
	for c := 0; c < m.num_cols; c++ {
		rv += "-"
	}
	rv += "+\n"
	for r := 0; r < m.num_rows; r++ {
		rv += "|"
		for c := 0; c < m.num_cols; c++ {
			rv += m.grid[r][c].String()
		}
		rv += "|\n"
	}
	rv += "+"
	for c := 0; c < m.num_cols; c++ {
		rv += "-"
	}
	rv += "+"
	return rv
}

func InitMaze(num_rows int, num_cols int, sparseness float64) Maze {
	m := Maze{num_rows, num_cols, nil}
	m.grid = make([][]Cell, num_rows)
	for r := 0; r < num_rows; r++ {
		m.grid[r] = make([]Cell, num_cols)
	}

	for r := 0; r < num_rows; r++ {
		for c := 0; c < num_cols; c++ {
			if rand.Float64() < sparseness {
				m.grid[r][c] = Blocked
			} else {
				m.grid[r][c] = Empty
			}
		}
	}
	return m
}

func (m Maze) SetGoal(goal MazeLocation) {
	m.grid[goal.row][goal.col] = Goal
}

func (m Maze) SetStart(start MazeLocation) {
	m.grid[start.row][start.col] = Start
}

func (m Maze) GoalTest(x MazeLocation) bool {
	return m.grid[x.row][x.col] == Goal
}

func (m Maze) Successors(x MazeLocation) []MazeLocation {
	var locations []MazeLocation
	delta := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, d := range delta {
		new_row := x.row + d[0]
		new_col := x.col + d[1]

		if new_row >= 0 && new_row < m.num_rows && new_col >= 0 && new_col < m.num_cols {
			if m.grid[new_row][new_col] != Blocked {
				locations = append(locations, MazeLocation{new_row, new_col})
			}
		}
	}

	return locations
}

func (m Maze) GetStart() MazeLocation {
	for r := 0; r < m.num_rows; r++ {
		for c := 0; c < m.num_cols; c++ {
			if m.grid[r][c] == Start {
				return MazeLocation{r, c}
			}
		}
	}
	return MazeLocation{-1, -1}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	num_rows := 10
	num_cols := 10
	m := InitMaze(num_rows, num_cols, .2)
	m.SetStart(MazeLocation{0, 0})
	m.SetGoal(MazeLocation{num_rows - 1, num_cols - 1})
	fmt.Println(m)

	fmt.Println(m.GoalTest(MazeLocation{5, 5}), m.GoalTest(MazeLocation{num_rows - 1, num_cols - 1}))

	start := m.GetStart()
	nbrs := m.Successors(start)
	fmt.Println(nbrs)
}
