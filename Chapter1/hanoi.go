package main

import (
	"classic_sc/stack"
	"fmt"
)

func hanoi_move(src *stack.Stack[int], dst *stack.Stack[int], scratch *stack.Stack[int], n int) {
	if n == 1 {
		val, _ := src.Pop()
		dst.Push(val)
	} else {
		hanoi_move(src, scratch, dst, n-1)
		hanoi_move(src, dst, scratch, 1)
		hanoi_move(scratch, dst, src, n-1)
	}
}

func main() {
	s := stack.New[int]()

	for i := 1; i <= 10; i++ {
		s.Push(i * i)
	}

	for {
		val, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(val)
	}

	num_discs := 3
	tower_a := stack.New[int]()
	tower_b := stack.New[int]()
	tower_c := stack.New[int]()

	for i := 1; i <= num_discs; i++ {
		tower_a.Push(i)
	}

	fmt.Println(tower_a, tower_b, tower_c)

	hanoi_move(tower_a, tower_b, tower_c, num_discs)

	fmt.Println(tower_a, tower_b, tower_c)
}
