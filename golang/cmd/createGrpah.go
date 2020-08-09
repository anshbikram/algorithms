package main

import (
	"algorithms/graph"
)

func main() {
	adjGraph := graph.CreateGraphFromStdin()
	adjGraph.Print()
}
