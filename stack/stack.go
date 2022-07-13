package stack

type Stack[T any] struct {
	vals     []T
	max_size int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
	if len(s.vals) > s.max_size {
		s.max_size = len(s.vals)
	}
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

func (s *Stack[T]) PopFirst() (T, bool) {
	var val T
	n := len(s.vals)
	if n > 0 {
		val = s.vals[0]
		s.vals = s.vals[1:]
		return val, true
	} else {
		return val, false
	}
}

func (s *Stack[T]) MaxSize() int {
	return s.max_size
}
