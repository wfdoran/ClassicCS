package graph

import (
	"classic_sc/stack"
	"fmt"
)

type Edge struct {
	src int
	dst int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d -> %d", e.src, e.dst)
}

func NewEdge(src int, dst int) Edge {
	return Edge{src: src, dst: dst}
}

func reverseEdge(e Edge) Edge {
	return Edge{src: e.dst, dst: e.src}
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

func (g *Graph[V]) AddEdgeByIndices(idx1 int, idx2 int) {
	e := NewEdge(idx1, idx2)
	g.AddEdge(e)
}

func (g *Graph[V]) AddEdgeByVertices(v1 V, v2 V) {
	idx1 := g.IndexOf(v1)
	idx2 := g.IndexOf(v2)
	g.AddEdgeByIndices(idx1, idx2)
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
		for _, v2 := range g.NeighborsForIndex(idx) {
			if first {
				first = false
			} else {
				rv += ", "
			}
			s := fmt.Sprint(v2)
			rv += s
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

func (n Node[T]) GetPath() []T {
	if n.parent == nil {
		return []T{n.state}
	} else {
		x := n.parent.GetPath()
		x = append(x, n.state)
		return x
	}
}

func (g Graph[V]) ShortestPath(src V, dst V) *Node[V] {
	startNode := Node[V]{state: src, parent: nil, cost: 0.0, heuristic: 0.0}

	frontier := stack.New[Node[V]]()
	frontier.Push(startNode)

	explored := make(map[V]bool)
	explored[src] = true

	for {
		curr, ok := frontier.PopFirst()
		if !ok {
			return nil
		}

		if curr.state == dst {
			return &curr
		}

		for _, nbr := range g.NeighborsForVertex(curr.state) {
			_, ok := explored[nbr]

			if !ok {
				nbrNode := Node[V]{nbr, &curr, curr.cost + 1.0, 0.0}
				frontier.Push(nbrNode)
				explored[nbr] = true
			}
		}
	}
}
