package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	expectedVertices := []*Vertex{
		{key: 1},
		{key: 2},
		{key: 3},
		{key: 4},
	}

	if !reflect.DeepEqual(g.vertices, expectedVertices) {
		t.Errorf("AddVertex: Expected vertices: %v, but got: %v", expectedVertices, g.vertices)
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	vertex1 := g.getVertex(1)
	vertex2 := g.getVertex(2)
	vertex3 := g.getVertex(3)

	if !g.contains(vertex1.adjacent, 2) || !g.contains(vertex1.adjacent, 3) {
		t.Errorf("AddEdge: Expected adjacent vertices of 1: %v, but got: [%v %v]", []int{2, 3}, vertex1.adjacent[0].key, vertex1.adjacent[1].key)
	}
	if vertex2.adjacent != nil || vertex3.adjacent != nil {
		t.Errorf("AddEdge: Unexpected adjacent vertices for vertex 2 or 3")
	}
}

func TestDFS(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)
	g.AddEdge(4, 1)

	result := make([]int, 0)
	g.DFS(g.getVertex(1), &result, make(map[int]bool))

	expectedResult := []int{1, 2, 4, 5, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("DFS: Expected traversal result: %v, but got: %v", expectedResult, result)
	}
}

func TestBFS(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)
	g.AddEdge(4, 1)

	result := make([]int, 0)
	g.BFS(g.getVertex(1), &result, make(map[int]bool))

	expectedResult := []int{1, 2, 4, 3, 5}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("BFS: Expected traversal result: %v, but got: %v", expectedResult, result)
	}
}
