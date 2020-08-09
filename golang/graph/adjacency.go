package graph

import "fmt"

// AdjacencyGraph is represented in adjanceny list
type AdjacencyGraph struct {
	vertices map[string][]string
	Directed bool
	Weighted bool
}

// AddEdge adds egde to the graph
func (graph *AdjacencyGraph) AddEdge(from, to string) {
	if graph.vertices == nil {
		graph.vertices = make(map[string][]string)
	}

	value := graph.vertices[from]
	value = append(value, to)
	graph.vertices[from] = value

	toValue := graph.vertices[to]
	if !graph.Directed {
		toValue = append(toValue, from)
	}
	graph.vertices[to] = toValue
}

// Print prints the graph
func (graph *AdjacencyGraph) Print() {
	if graph.vertices == nil {
		fmt.Println("Graph is empty")
		return
	}

	fmt.Println("### Printing Graph Start ###")
	for key, value := range graph.vertices {
		fmt.Println(key, value)
	}
	fmt.Println("### Printing Graph End ###")
}

// GetAdjacencyList get the adjacency representation
func (graph *AdjacencyGraph) GetAdjacencyList() map[string][]string {
	return graph.vertices
}

// CreateGraphFromStdin creates graph by input from stdin
func CreateGraphFromStdin() AdjacencyGraph {
	fmt.Println("Is the graph directed (1) or undirected (0)")
	var isDir int
	fmt.Scan(&isDir)
	fmt.Println("Is the graph weighted (1) or not weighted (0)")
	var isWgt int
	fmt.Scan(&isDir)
	graph := AdjacencyGraph{Weighted: isWgt == 1}

	fmt.Println("Please enter edges <source> <destination>. enter `-1 -1` to end input")
	for true {
		var x, y string
		fmt.Scan(&x, &y)
		if x == "-1" {
			break
		}

		graph.AddEdge(x, y)
	}

	return graph
}
