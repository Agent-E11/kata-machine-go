package dijkstra_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/src/dijkstra"
	"github.com/agent-e11/kata-machine-go/tests"
)

func TestDijkstraList(t *testing.T) {
	path := dijkstra.DijkstraList(0, 6, tests.GraphList1)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
