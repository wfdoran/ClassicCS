package main

import (
	"classic_sc/stack"
	"fmt"
)

const MAX_NUM int = 3

const (
	WEST = 0
	EAST = 1
)

type State struct {
	miss [2]int
	cann [2]int
	boat int
}

func (s State) String() string {
	rv := ""
	for i := 0; i < s.miss[WEST]; i++ {
		rv += "M"
	}
	for i := 0; i < s.cann[WEST]; i++ {
		rv += "X"
	}

	if s.boat == WEST {
		rv += " w_____  "
	} else {
		rv += "  _____w "
	}
	for i := 0; i < s.cann[EAST]; i++ {
		rv += "X"
	}
	for i := 0; i < s.miss[EAST]; i++ {
		rv += "M"
	}
	return rv
}

func (s State) IsLegal() bool {
	if s.miss[EAST] > 0 && s.miss[EAST] < s.cann[EAST] {
		return false
	}
	if s.miss[WEST] > 0 && s.miss[WEST] < s.cann[WEST] {
		return false
	}
	return true
}

func (s State) Successors() []State {
	var sucs []State

	curr := s.boat
	other := 1 - curr

	if s.miss[curr] >= 2 {
		x := s
		x.miss[curr] -= 2
		x.miss[other] += 2
		x.boat = other
		if x.IsLegal() {
			sucs = append(sucs, x)
		}
	}
	if s.miss[curr] >= 1 {
		x := s
		x.miss[curr] -= 1
		x.miss[other] += 1
		x.boat = other
		if x.IsLegal() {
			sucs = append(sucs, x)
		}
	}
	if s.cann[curr] >= 2 {
		x := s
		x.cann[curr] -= 2
		x.cann[other] += 2
		x.boat = other
		if x.IsLegal() {
			sucs = append(sucs, x)
		}
	}
	if s.cann[curr] >= 1 {
		x := s
		x.cann[curr] -= 1
		x.cann[other] += 1
		x.boat = other
		if x.IsLegal() {
			sucs = append(sucs, x)
		}
	}
	if s.miss[curr] >= 1 && s.cann[curr] >= 1 {
		x := s
		x.miss[curr] -= 1
		x.miss[other] += 1
		x.cann[curr] -= 1
		x.cann[other] += 1
		x.boat = other
		if x.IsLegal() {
			sucs = append(sucs, x)
		}
	}

	return sucs
}

func (s State) GoalTest() bool {
	return s.miss[WEST] == 0 && s.cann[WEST] == 0
}

type Node[T any] struct {
	state     T
	parent    *Node[T]
	cost      float64
	heuristic float64
}

func bfs(start State) *Node[State] {
	startNode := Node[State]{start, nil, 0.0, 0.0}

	frontier := stack.New[Node[State]]()
	frontier.Push(startNode)

	explored := make(map[State]bool)
	explored[start] = true

	for {
		curr, ok := frontier.PopFirst()
		if !ok {
			return nil
		}

		if curr.state.GoalTest() {
			return &curr
		}

		for _, nbr := range curr.state.Successors() {
			_, ok := explored[nbr]

			if !ok {
				nbrNode := Node[State]{nbr, &curr, curr.cost + 1.0, 0.0}
				frontier.Push(nbrNode)
				explored[nbr] = true
			}
		}
	}
}

func main() {
	start := State{miss: [2]int{MAX_NUM, 0}, cann: [2]int{MAX_NUM, 0}, boat: WEST}
	fmt.Println(start)

	path := bfs(start)

	for path != nil {
		fmt.Println(path.state)
		path = path.parent
	}
}
