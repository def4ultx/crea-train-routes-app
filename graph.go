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

func NewGraph() *Graph {
	return &Graph{
		Size:    0,
		Mapping: make(map[string]int),
		Nodes:   make([][]int, 0),
	}
}

// Add vertex to current graph
func (g *Graph) AddVertex(name string) {
	_, ok := g.Mapping[name]
	if ok {
		return
	}

	g.Size++
	g.Mapping[name] = g.Size
}

// Add edge to current graph
func (g *Graph) AddEdge(src, dest string, dist int) {

}

// Calculate shortest distance with number of hop between 2 vertex
func (g *Graph) ShortestPath(src, dest string) (hop, dist int) {
	return 0, 0
}

// Represent single vertex
type Vertex struct {
}
