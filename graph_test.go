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

	assert.Equal(t, graph.Size, 4)
	assert.Equal(t, graph.Mapping["A"], 1)
	assert.Equal(t, graph.Mapping["B"], 2)
	assert.Equal(t, graph.Mapping["C"], 3)
	assert.Equal(t, graph.Mapping["D"], 4)
}

func TestGraph_AddVertex_HaveDuplicateVertex(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")

	assert.Equal(t, graph.Size, 3)
	assert.Equal(t, graph.Mapping["A"], 1)
	assert.Equal(t, graph.Mapping["B"], 2)
	assert.Equal(t, graph.Mapping["C"], 3)
}
