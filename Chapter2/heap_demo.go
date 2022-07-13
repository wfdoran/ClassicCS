package main

import (
	"classic_sc/heap"
	"fmt"
	"math/rand"
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

	for i := 0; i < 20; i++ {
		p := Pair{rand.Intn(100), rand.Intn(100)}
		h.Push(p)
	}

	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		fmt.Println(v, v.Score())
	}
}
