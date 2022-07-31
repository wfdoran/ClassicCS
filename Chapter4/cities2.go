package main

import (
	"classic_sc/wt_graph"
	"fmt"
)

func main() {
	g := wt_graph.New[string]()

	g.AddVertex("Seattle")
	g.AddVertex("San Francisco")
	g.AddVertex("Los Angeles")
	g.AddVertex("Riverside")
	g.AddVertex("Phoenix")
	g.AddVertex("Boston")
	g.AddVertex("Atlanta")
	g.AddVertex("Miami")
	g.AddVertex("Dallas")
	g.AddVertex("Philadelphia")
	g.AddVertex("Detroit")
	g.AddVertex("Chicago")
	g.AddVertex("New York")
	g.AddVertex("Houston")
	g.AddVertex("Washington")

	g.AddEdgeByVertices("Seattle", "Chicago", 1737.0)
	g.AddEdgeByVertices("Seattle", "San Francisco", 678.0)
	g.AddEdgeByVertices("San Francisco", "Riverside", 386.0)
	g.AddEdgeByVertices("San Francisco", "Los Angeles", 348.0)
	g.AddEdgeByVertices("Los Angeles", "Riverside", 50.0)
	g.AddEdgeByVertices("Los Angeles", "Phoenix", 357.0)
	g.AddEdgeByVertices("Riverside", "Phoenix", 307.0)
	g.AddEdgeByVertices("Riverside", "Chicago", 1704.0)
	g.AddEdgeByVertices("Phoenix", "Dallas", 887.0)
	g.AddEdgeByVertices("Phoenix", "Houston", 1015.0)
	g.AddEdgeByVertices("Dallas", "Chicago", 805)
	g.AddEdgeByVertices("Dallas", "Houston", 225.0)
	g.AddEdgeByVertices("Dallas", "Atlanta", 721.0)
	g.AddEdgeByVertices("Houston", "Atlanta", 702.0)
	g.AddEdgeByVertices("Houston", "Miami", 968.0)
	g.AddEdgeByVertices("Atlanta", "Chicago", 588.0)
	g.AddEdgeByVertices("Atlanta", "Washington", 543.0)
	g.AddEdgeByVertices("Atlanta", "Miami", 604.0)
	g.AddEdgeByVertices("Miami", "Washington", 923.0)
	g.AddEdgeByVertices("Chicago", "Detroit", 238.0)
	g.AddEdgeByVertices("Detroit", "Boston", 613.0)
	g.AddEdgeByVertices("Detroit", "Washington", 396.0)
	g.AddEdgeByVertices("Detroit", "New York", 396.0)
	g.AddEdgeByVertices("Boston", "New York", 190.0)
	g.AddEdgeByVertices("New York", "Philadelphia", 81.0)
	g.AddEdgeByVertices("Philadelphia", "Washington", 123.0)

	fmt.Println(g)

	n := g.ShortestPath("Boston", "Miami")
	path := n.GetPath()
	fmt.Println(path)
}
