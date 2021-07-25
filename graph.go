package main

func createGraphFromFile(filepath string) (*Graph, error) {
	return nil, nil
}

// Hold all nodes, using adjacency metric representation
type Graph struct {
	Size    int
	Mapping map[string]int
	Nodes   [][]int
}

// Add vertex to current graph
func (g *Graph) AddVertex(name string) {

}

// Add edge to current graph
func (g *Graph) AddEdge(src, dest string, dist int) {

}

// Calculate shortest path with number of hop between 2 vertex
func (g *Graph) ShortestPath(src, dest string) (hop, dist int) {
	return 0, 0
}

// Represent single vertex
type Vertex struct {
}
