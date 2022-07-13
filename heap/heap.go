package heap

type heaper interface {
	Score() float64
}

type Heap[T heaper] struct {
	vals     []T
	max_size int
}

func New[T heaper]() *Heap[T] {
	return &Heap[T]{nil, 0}
}

func (h *Heap[T]) Push(val T) {
	h.vals = append(h.vals, val)
	if len(h.vals) > h.max_size {
		h.max_size = len(h.vals)
	}

	idx := len(h.vals) - 1

	for idx > 0 {
		parent := (idx - 1) / 2

		if h.vals[idx].Score() > h.vals[parent].Score() {
			h.vals[idx], h.vals[parent] = h.vals[parent], h.vals[idx]
			idx = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	var val T
	n := len(h.vals)
	if n == 0 {
		return val, false
	}

	val = h.vals[0]
	if n > 1 {
		h.vals[0] = h.vals[n-1]
	}
	h.vals = h.vals[:n-1]
	m := len(h.vals)
	idx := 0
	for {
		child0 := 2*idx + 1
		child1 := 2*idx + 2

		if child0 >= m {
			break
		}

		if child1 >= m {
			if h.vals[idx].Score() < h.vals[child0].Score() {
				h.vals[idx], h.vals[child0] = h.vals[child0], h.vals[idx]
			}
			break
		}

		if h.vals[idx].Score() >= h.vals[child0].Score() && h.vals[idx].Score() > h.vals[child1].Score() {
			break
		}

		if h.vals[child0].Score() > h.vals[child1].Score() {
			h.vals[idx], h.vals[child0] = h.vals[child0], h.vals[idx]
			idx = child0
		} else {
			h.vals[idx], h.vals[child1] = h.vals[child1], h.vals[idx]
			idx = child1
		}
	}
	return val, true
}

func (h *Heap[T]) MaxSize() int {
	return h.max_size
}
