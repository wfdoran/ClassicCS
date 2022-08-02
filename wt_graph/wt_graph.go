package wt_graph

import (
	"classic_sc/heap"
	"fmt"
)

type Edge struct {
	src    int
	dst    int
	weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%d -> %d (%f)", e.src, e.dst, e.weight)
}

func NewEdge(src int, dst int, weight float64) Edge {
	return Edge{src: src, dst: dst, weight: weight}
}

func reverseEdge(e Edge) Edge {
	return Edge{src: e.dst, dst: e.src, weight: e.weight}
}

type Graph[V comparable] struct {
	vertices []V
	edges    [][]Edge
}

func New[V comparable]() *Graph[V] {
	var vertices []V
	var edges [][]Edge
	return &Graph[V]{vertices: vertices, edges: edges}
}

func (g Graph[V]) VertexCount() int {
	return len(g.vertices)
}

func (g Graph[V]) EdgeCount() int {
	sum := 0
	for _, nbrs := range g.edges {
		sum += len(nbrs)
	}
	return sum
}

func (g *Graph[V]) AddVertex(v V) int {
	g.vertices = append(g.vertices, v)
	var nbrs []Edge
	g.edges = append(g.edges, nbrs)
	return g.VertexCount() - 1
}

func (g *Graph[V]) AddEdge(e Edge) {
	g.edges[e.src] = append(g.edges[e.src], e)
	g.edges[e.dst] = append(g.edges[e.dst], reverseEdge(e))
}

func (g Graph[V]) VertexAt(idx int) V {
	return g.vertices[idx]
}

func (g Graph[V]) IndexOf(v V) int {
	for idx, vv := range g.vertices {
		if v == vv {
			return idx
		}
	}
	fmt.Println("Fail on ", v)
	return -1
}

func (g *Graph[V]) AddEdgeByIndices(idx1 int, idx2 int, weight float64) {
	e := NewEdge(idx1, idx2, weight)
	g.AddEdge(e)
}

func (g *Graph[V]) AddEdgeByVertices(v1 V, v2 V, weight float64) {
	idx1 := g.IndexOf(v1)
	idx2 := g.IndexOf(v2)
	g.AddEdgeByIndices(idx1, idx2, weight)
}

func (g Graph[V]) NeighborsForIndex(idx int) []V {
	var rv []V
	for _, e := range g.edges[idx] {
		rv = append(rv, g.vertices[e.dst])
	}
	return rv
}

func (g Graph[V]) NeighborsForVertex(v V) []V {
	idx := g.IndexOf(v)
	return g.NeighborsForIndex(idx)
}

func (g Graph[V]) EdgesForIndex(idx int) []Edge {
	return g.edges[idx]
}

func (g Graph[V]) EdgesForVertex(v V) []Edge {
	idx := g.IndexOf(v)
	return g.EdgesForIndex(idx)
}

func (g Graph[V]) String() string {
	rv := ""

	for idx, v := range g.vertices {
		s := fmt.Sprint(v)
		rv += s
		rv += " -> ["
		first := true
		for _, nbr := range g.edges[idx] {
			if first {
				first = false
			} else {
				rv += ", "
			}
			s := fmt.Sprint(g.vertices[nbr.dst])
			rv += s
			s1 := fmt.Sprintf("(%.1f)", nbr.weight)
			rv += s1
		}
		rv += "]\n"
	}

	return rv
}

type Node[T any] struct {
	state     T
	parent    *Node[T]
	cost      float64
	heuristic float64
}

func (n Node[T]) GetCost() float64 {
	return n.cost
}

func (n Node[T]) GetPath() []T {
	if n.parent == nil {
		return []T{n.state}
	} else {
		x := n.parent.GetPath()
		x = append(x, n.state)
		return x
	}
}

func (n Node[T]) Score() float64 {
	return -n.cost
}

func (g Graph[V]) ShortestPath(src V, dst V) *Node[V] {
	startNode := Node[V]{state: src, parent: nil, cost: 0.0, heuristic: 0.0}

	frontier := heap.New[Node[V]]()
	frontier.Push(startNode)

	explored := make(map[V]float64)
	explored[src] = 0.0

	var rv *Node[V] = nil
	for {
		curr, ok := frontier.Pop()
		if !ok {
			return rv
		}

		if curr.state == dst {
			if rv == nil || curr.cost < rv.cost {
				rv = &curr
			}
		} else {
			idx := g.IndexOf(curr.state)
			for _, e := range g.edges[idx] {
				nbr := g.VertexAt(e.dst)
				nbr_cost := curr.cost + e.weight

				cost, ok := explored[nbr]

				if !ok || nbr_cost < cost {
					nbrNode := Node[V]{nbr, &curr, nbr_cost, 0.0}
					frontier.Push(nbrNode)
					explored[nbr] = nbr_cost
				}
			}
		}
	}
}

func (e Edge) Score() float64 {
	return -e.weight
}

func PathWeight(path []Edge) float64 {
	total := 0.0
	for _, e := range path {
		total += e.weight
	}
	return total
}

func (g Graph[V]) Mst(start int) []Edge {
	var rv []Edge
	pq := heap.New[Edge]()
	visited := make([]bool, g.VertexCount())

	visited[start] = true
	for _, e := range g.edges[start] {
		pq.Push(e)
	}

	for {
		e, ok := pq.Pop()
		if !ok {
			break
		}

		if visited[e.dst] {
			continue
		}
		visited[e.dst] = true
		for _, e2 := range g.edges[e.dst] {
			pq.Push(e2)
		}
		rv = append(rv, e)
	}

	return rv
}
