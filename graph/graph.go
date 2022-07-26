package graph

import "fmt"

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

type Vertexable interface {
	comparable
	fmt.Stringer
}

type Graph[V Vertexable] struct {
	vertices []V
	edges    [][]Edge
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

func (g Graph[V]) AddVertex(v V) int {
	g.vertices = append(g.vertices, v)
	var nbrs []Edge
	g.edges = append(g.edges, nbrs)
	return g.VertexCount() - 1
}

func (g Graph[V]) AddEdge(e Edge) {
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
	return -1
}

func (g Graph[V]) AddEdgeByIndices(idx1 int, idx2 int) {
	e := NewEdge(idx1, idx2)
	g.AddEdge(e)
}

func (g Graph[V]) AddEdgeByVertices(v1 V, v2 V) {
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
		rv += v.String()
		rv += " -> ["
		first := true
		for _, v2 := range g.NeighborsForIndex(idx) {
			if first {
				first = false
			} else {
				rv += ", "
			}
			rv += v2.String()
		}
		rv += " ]\n"
	}

	return rv
}
