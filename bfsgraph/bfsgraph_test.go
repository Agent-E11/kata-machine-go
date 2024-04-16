package bfsgraph_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/bfsgraph"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestBFSGraphList(t *testing.T) {
	path := bfsgraph.BFSList(graph.List2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}

	path = bfsgraph.BFSList(graph.List2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}

func TestBFSGraphMatrix(t *testing.T) {
	path := bfsgraph.BFSMatrix(graph.Matrix2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}
	
	if !slices.Equal(path, expect) {
		t.Fatalf("expected path from 0 to 6 to be %v, got %v", expect, path)
	}

	path = bfsgraph.BFSMatrix(graph.Matrix2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path from 6 to 0 to be %v, got %v", expect, path)
	}
}
