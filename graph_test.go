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
