package main

import (
	"classic_sc/stack"
	"fmt"
)

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
}
