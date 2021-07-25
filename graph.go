package main

import (
	"math"
)

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

	nodes := make([]int, g.Size-1)
	g.Nodes = append(g.Nodes, nodes)
	for i := range g.Nodes {
		g.Nodes[i] = append(g.Nodes[i], 0)
	}
}

// Add edge to current graph
func (g *Graph) AddEdge(src, dest string, dist int) {
	srcIndex, ok := g.Mapping[src]
	if !ok {
		g.AddVertex(src)
		srcIndex = g.Mapping[src]
	}

	destIndex, ok := g.Mapping[dest]
	if !ok {
		g.AddVertex(dest)
		destIndex = g.Mapping[dest]
	}

	g.Nodes[srcIndex-1][destIndex-1] = dist
	g.Nodes[destIndex-1][srcIndex-1] = dist
}

// Calculate shortest distance with number of hop between 2 vertex
func (g *Graph) ShortestPath(src, dest string) (hop, distance int) {
	srcMapping, ok := g.Mapping[src]
	if !ok {
		return -1, -1
	}
	destMapping, ok := g.Mapping[dest]
	if !ok {
		return -1, -1
	}

	var (
		srcIndex  = srcMapping - 1
		destIndex = destMapping - 1

		dist = make([]int, g.Size)
		seen = make([]bool, g.Size)
		prev = make([]int, g.Size)
	)
	for i := 0; i < g.Size; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[srcIndex] = 0

	for {
		u := findMin(dist, seen)
		if u == -1 || u == destIndex {
			break
		}
		seen[u] = true

		for i, v := range g.Nodes[u] {
			if v == 0 {
				continue
			}

			alt := dist[u] + v
			if alt < dist[i] {
				dist[i] = alt
				prev[i] = u
			}
		}
	}

	if dist[destIndex] == math.MaxInt32 {
		return -1, -1
	}

	return calculateHop(destIndex, prev), dist[destIndex]
}

func calculateHop(start int, prev []int) int {
	var hop int
	index := start
	for {
		i := prev[index]
		if i == -1 {
			break
		}

		hop++
		index = i
	}
	return hop
}

func findMin(nodes []int, seen []bool) int {
	index := -1
	current := math.MaxInt32

	for i, v := range nodes {
		if v < current && !seen[i] {
			current = v
			index = i
		}
	}
	return index
}
