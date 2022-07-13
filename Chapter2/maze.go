package main

import (
	"classic_sc/heap"
	"classic_sc/stack"
	"fmt"
	"math"
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

func (m Maze) GetGoal() MazeLocation {
	for r := 0; r < m.num_rows; r++ {
		for c := 0; c < m.num_cols; c++ {
			if m.grid[r][c] == Goal {
				return MazeLocation{r, c}
			}
		}
	}
	return MazeLocation{-1, -1}
}

type Node[T any] struct {
	state     T
	parent    *Node[T]
	cost      float64
	heuristic float64
}

func EucleanDistance(a MazeLocation, b MazeLocation) float64 {
	delta_row := float64(a.row - b.row)
	delta_col := float64(a.col - b.col)
	return math.Sqrt(delta_row*delta_row + delta_col*delta_col)
}

func ManhattanDistance(a MazeLocation, b MazeLocation) float64 {
	delta_row := float64(a.row - b.row)
	delta_col := float64(a.col - b.col)
	return math.Abs(delta_row) + math.Abs(delta_col)
}

func (node Node[MazeLocation]) Score() float64 {
	return -(node.cost + node.heuristic)
}

func (m Maze) dfs() *Node[MazeLocation] {
	start := m.GetStart()
	startNode := Node[MazeLocation]{start, nil, 0.0, 0.0}

	frontier := stack.New[Node[MazeLocation]]()
	frontier.Push(startNode)

	explored := make(map[MazeLocation]bool)
	explored[start] = true

	for {
		curr, ok := frontier.Pop()
		if !ok {
			return nil
		}

		if m.GoalTest(curr.state) {
			return &curr
		}

		for _, nbr := range m.Successors(curr.state) {
			_, ok := explored[nbr]

			if !ok {
				nbrNode := Node[MazeLocation]{nbr, &curr, curr.cost + 1.0, 0.0}
				frontier.Push(nbrNode)
				explored[nbr] = true
			}
		}
	}
}

func (m Maze) bfs() *Node[MazeLocation] {
	start := m.GetStart()
	startNode := Node[MazeLocation]{start, nil, 0.0, 0.0}

	frontier := stack.New[Node[MazeLocation]]()
	frontier.Push(startNode)

	explored := make(map[MazeLocation]bool)
	explored[start] = true

	for {
		curr, ok := frontier.PopFirst()
		if !ok {
			return nil
		}

		if m.GoalTest(curr.state) {
			return &curr
		}

		for _, nbr := range m.Successors(curr.state) {
			_, ok := explored[nbr]

			if !ok {
				nbrNode := Node[MazeLocation]{nbr, &curr, curr.cost + 1.0, 0.0}
				frontier.Push(nbrNode)
				explored[nbr] = true
			}
		}
	}
}

func (m Maze) a_star() *Node[MazeLocation] {
	// dist_func := ManhattanDistance
	dist_func := EucleanDistance
	goal := m.GetGoal()

	start := m.GetStart()
	startNode := Node[MazeLocation]{start, nil, 0.0, dist_func(start, goal)}

	frontier := heap.New[Node[MazeLocation]]()
	frontier.Push(startNode)

	explored := make(map[MazeLocation]bool)
	explored[start] = true

	for {
		curr, ok := frontier.Pop()
		if !ok {
			return nil
		}

		if m.GoalTest(curr.state) {
			return &curr
		}

		for _, nbr := range m.Successors(curr.state) {
			_, ok := explored[nbr]

			if !ok {
				nbrNode := Node[MazeLocation]{nbr, &curr, curr.cost + 1.0, dist_func(nbr, goal)}
				frontier.Push(nbrNode)
				explored[nbr] = true
			}
		}
	}
}

func (m *Maze) ClearPath() {
	for r := 0; r < m.num_rows; r++ {
		for c := 0; c < m.num_cols; c++ {
			if m.grid[r][c] == Path {
				m.grid[r][c] = Empty
			}
		}
	}
}

func (m *Maze) MarkPath(node *Node[MazeLocation]) {
	curr := node

	for curr != nil {
		r := curr.state.row
		c := curr.state.col

		if m.grid[r][c] == Empty {
			m.grid[r][c] = Path
		}
		curr = curr.parent
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	num_rows := 10
	num_cols := 10
	sparseness := .2
	m := InitMaze(num_rows, num_cols, sparseness)
	m.SetStart(MazeLocation{0, 0})
	m.SetGoal(MazeLocation{num_rows - 1, num_cols - 1})
	fmt.Println(m)

	fmt.Println(m.GoalTest(MazeLocation{5, 5}), m.GoalTest(MazeLocation{num_rows - 1, num_cols - 1}))

	start := m.GetStart()
	nbrs := m.Successors(start)
	fmt.Println(nbrs)

	t := m.dfs()

	if t != nil {
		fmt.Println("cost = ", t.cost)
		m.MarkPath(t)
		fmt.Println(m)
	}
	m.ClearPath()

	t = m.bfs()

	if t != nil {
		fmt.Println("cost = ", t.cost)
		m.MarkPath(t)
		fmt.Println(m)
	}
	m.ClearPath()

	t = m.a_star()

	if t != nil {
		fmt.Println("cost = ", t.cost)
		m.MarkPath(t)
		fmt.Println(m)
	}
}
