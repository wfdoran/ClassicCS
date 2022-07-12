package main

import (
	"classic_sc/heap"
	"fmt"
)

type Pair struct {
	a int
	b int
}

func (p Pair) Score() float64 {
	aa := float64(p.a)
	bb := float64(p.b)
	return 1.499*aa + bb
}

func main() {
	h := heap.New[Pair]()

	h.Push(Pair{1, 2})
	h.Push(Pair{2, 1})
	h.Push(Pair{4, 0})
	h.Push(Pair{1, 1})

	fmt.Println(h)

	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		fmt.Println(v, v.Score())
	}

}
