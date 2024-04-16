package dijkstra_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/dijkstra"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestDijkstraList(t *testing.T) {
	path := dijkstra.DijkstraList(0, 6, graph.List1)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
