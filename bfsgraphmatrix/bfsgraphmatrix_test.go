package bfsgraphmatrix_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/bfsgraphmatrix"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestBFSGraphMatrix(t *testing.T) {
	path := bfsgraphmatrix.BFS(graph.Matrix2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}
	
	if !slices.Equal(path, expect) {
		t.Fatalf("expected path from 0 to 6 to be %v, got %v", expect, path)
	}

	path = bfsgraphmatrix.BFS(graph.Matrix2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path from 6 to 0 to be %v, got %v", expect, path)
	}
}
