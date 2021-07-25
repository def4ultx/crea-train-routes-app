package main

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	actual := NewGraph()
	expected := &Graph{
		Size:    0,
		Mapping: map[string]int{},
		Nodes:   [][]int{},
	}
	assert.Equal(t, expected, actual)
}

func TestGraph_AddVertex_AddNewVertex(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")

	assert.Equal(t, 4, graph.Size)
	assert.Equal(t, 1, graph.Mapping["A"])
	assert.Equal(t, 2, graph.Mapping["B"])
	assert.Equal(t, 3, graph.Mapping["C"])
	assert.Equal(t, 4, graph.Mapping["D"])

	expected := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddVertex_HaveDuplicateVertex(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")

	assert.Equal(t, 3, graph.Size)
	assert.Equal(t, 1, graph.Mapping["A"])
	assert.Equal(t, 2, graph.Mapping["B"])
	assert.Equal(t, 3, graph.Mapping["C"])

	expected := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddVertexToExistingGraph(t *testing.T) {
	graph := &Graph{
		Size: 3,
		Mapping: map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
		Nodes: [][]int{
			{0, 1, 2},
			{1, 0, 2},
			{2, 2, 0},
		},
	}

	graph.AddVertex("D")
	expected := [][]int{
		{0, 1, 2, 0},
		{1, 0, 2, 0},
		{2, 2, 0, 0},
		{0, 0, 0, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddEdge_ExistingVertex(t *testing.T) {
	graph := &Graph{
		Size: 3,
		Mapping: map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
		Nodes: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 2)

	expected := [][]int{
		{0, 1, 2},
		{1, 0, 0},
		{2, 0, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddEdge_1_ExistVertex(t *testing.T) {
	graph := &Graph{
		Size: 3,
		Mapping: map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
		Nodes: [][]int{
			{0, 1, 2},
			{1, 0, 0},
			{2, 0, 0},
		},
	}

	graph.AddEdge("C", "D", 5)
	expected := [][]int{
		{0, 1, 2, 0},
		{1, 0, 0, 0},
		{2, 0, 0, 5},
		{0, 0, 5, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddEdge_NewVertex(t *testing.T) {
	graph := &Graph{
		Size: 3,
		Mapping: map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
		Nodes: [][]int{
			{0, 1, 2},
			{1, 0, 0},
			{2, 0, 0},
		},
	}

	graph.AddEdge("D", "E", 5)
	expected := [][]int{
		{0, 1, 2, 0, 0},
		{1, 0, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 5},
		{0, 0, 0, 5, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func TestGraph_AddEdge_NewGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge("A", "B", 1)

	expected := [][]int{
		{0, 1},
		{1, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}

func newTestGraph() *Graph {
	graph := NewGraph()
	graph.AddEdge("A", "B", 5)
	graph.AddEdge("B", "C", 5)
	graph.AddEdge("C", "D", 7)
	graph.AddEdge("A", "D", 15)
	graph.AddEdge("E", "F", 5)
	graph.AddEdge("F", "G", 5)
	graph.AddEdge("G", "H", 10)
	graph.AddEdge("H", "I", 10)
	graph.AddEdge("I", "J", 5)
	graph.AddEdge("G", "J", 20)
	return graph
}

func TestGraph_ShortestPath_SourceNotExist(t *testing.T) {
	graph := newTestGraph()
	hop, dist := graph.ShortestPath("W", "A")
	assert.Equal(t, -1, hop)
	assert.Equal(t, -1, dist)
}

func TestGraph_ShortestPath_DestNotExist(t *testing.T) {
	graph := newTestGraph()
	hop, dist := graph.ShortestPath("A", "W")
	assert.Equal(t, -1, hop)
	assert.Equal(t, -1, dist)
}

func TestGraph_ShortestPath(t *testing.T) {
	testcases := []struct {
		src  string
		dest string

		hop  int
		dist int
	}{
		{"A", "B", 1, 5},
		{"A", "C", 2, 10},
		{"E", "J", 3, 30},
		{"A", "D", 1, 15},
		{"A", "J", -1, -1},
		{"A", "A", 0, 0},
		{"E", "I", 4, 30},
		{"B", "D", 2, 12},
	}

	for _, v := range testcases {
		t.Run("shortest path", func(t *testing.T) {
			graph := newTestGraph()
			hop, dist := graph.ShortestPath(v.src, v.dest)

			assert.Equal(t, v.hop, hop)
			assert.Equal(t, v.dist, dist)
		})
	}
}

func TestCreateGraphFromReader(t *testing.T) {
	csv := "A,B,5\nB,C,5\nC,D,7\nA,D,15\nE,F,5\nF,G,5\nG,H,10\nH,I,10\nI,J,5\nG,J,20\n"
	data := strings.NewReader(csv)

	actual, err := createGraphFromReader(data)
	graph := newTestGraph()

	assert.NoError(t, err)
	assert.Equal(t, graph, actual)
}

func TestCreateGraphFromReader_InvalidFormat(t *testing.T) {
	csv := "A,B,C"
	data := strings.NewReader(csv)

	actual, err := createGraphFromReader(data)
	assert.Error(t, err)
	assert.Empty(t, actual)
}

func TestCreateGraphFromReader_ReaderError(t *testing.T) {
	r := iotest.ErrReader(errors.New("custom error"))
	actual, err := createGraphFromReader(r)
	assert.Error(t, err)
	assert.Empty(t, actual)
}
