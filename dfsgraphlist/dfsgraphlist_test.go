package dfsgraphlist_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/dfsgraphlist"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestDFS(t *testing.T) {
	path := dfsgraphlist.DFS(graph.List2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}

	path = dfsgraphlist.DFS(graph.List2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
