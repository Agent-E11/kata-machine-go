package bfsgraphlist_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/bfsgraphlist"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestBFSGraphList(t *testing.T) {
	path := bfsgraphlist.BFS(graph.List2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}

	path = bfsgraphlist.BFS(graph.List2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
