package dijkstralist_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/dijkstralist"
	"github.com/agent-e11/kata-machine-go/graph"
)

func TestDijkstraList(t *testing.T) {
	path := dijkstralist.DijkstraList(0, 6, graph.List1)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
