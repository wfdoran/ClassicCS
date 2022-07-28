package main

import (
	"classic_sc/graph"
	"fmt"
)

func main() {
	g := graph.New[string]()

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

	g.AddEdgeByVertices("Seattle", "Chicago")
	g.AddEdgeByVertices("Seattle", "San Francisco")
	g.AddEdgeByVertices("San Francisco", "Riverside")
	g.AddEdgeByVertices("San Francisco", "Los Angeles")
	g.AddEdgeByVertices("Los Angeles", "Riverside")
	g.AddEdgeByVertices("Los Angeles", "Phoenix")
	g.AddEdgeByVertices("Riverside", "Phoenix")
	g.AddEdgeByVertices("Riverside", "Chicago")
	g.AddEdgeByVertices("Phoenix", "Dallas")
	g.AddEdgeByVertices("Phoenix", "Houston")
	g.AddEdgeByVertices("Dallas", "Chicago")
	g.AddEdgeByVertices("Dallas", "Houston")
	g.AddEdgeByVertices("Dallas", "Atlanta")
	g.AddEdgeByVertices("Houston", "Atlanta")
	g.AddEdgeByVertices("Houston", "Miami")
	g.AddEdgeByVertices("Atlanta", "Chicago")
	g.AddEdgeByVertices("Atlanta", "Washington")
	g.AddEdgeByVertices("Atlanta", "Miami")
	g.AddEdgeByVertices("Miami", "Washington")
	g.AddEdgeByVertices("Chicago", "Detroit")
	g.AddEdgeByVertices("Detroit", "Boston")
	g.AddEdgeByVertices("Detroit", "Washington")
	g.AddEdgeByVertices("Detroit", "New York")
	g.AddEdgeByVertices("Boston", "New York")
	g.AddEdgeByVertices("New York", "Philadelphia")
	g.AddEdgeByVertices("Philadelphia", "Washington")

	fmt.Println(g)

	n := g.ShortestPath("Boston", "Miami")
	path := n.GetPath()
	fmt.Println(path)
}
