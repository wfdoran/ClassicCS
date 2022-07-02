package stack

type Stack[T any] struct {
	vals []T
}

func New[T any]() *Stack[T] {
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
