package main

import (
	"testing"

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
			{0, 0, 2},
			{2, 2, 0},
		},
	}

	graph.AddVertex("D")
	expected := [][]int{
		{0, 1, 2, 0},
		{0, 0, 2, 0},
		{2, 2, 0, 0},
		{0, 0, 0, 0},
	}
	assert.Equal(t, expected, graph.Nodes)
}
