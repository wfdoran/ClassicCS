package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	var val T
	n := len(s.vals)
	if n > 0 {
		val = s.vals[n-1]
		s.vals = s.vals[:n-1]
		return val, true
	} else {
		return val, false
	}

}

func main() {
	s := NewStack[int]()

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
