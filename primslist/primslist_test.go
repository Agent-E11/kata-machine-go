package primslist_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/graph"
	"github.com/agent-e11/kata-machine-go/primslist"
)

func TestPrimsList(t *testing.T) {
	result := primslist.Prims(graph.List1)
	expect := graph.WeightedAdjacencyList{
        {
            { To: 2, Weight: 1 },
            { To: 1, Weight: 3 },
        },
        {
            { To: 0, Weight: 3 },
            { To: 4, Weight: 1 },
        },
        {
			{ To: 0, Weight: 1 },
		},
        {
			{ To: 6, Weight: 1 },
		},
        {
            { To: 1, Weight: 1 },
            { To: 5, Weight: 2 },
        },
        {
            { To: 4, Weight: 2 },
            { To: 6, Weight: 1 },
        },
        {
            { To: 5, Weight: 1 },
            { To: 3, Weight: 1 },
        },
	}

	if !slices.EqualFunc(
		result, expect,
		slices.Equal[[]graph.GraphEdge, graph.GraphEdge],
	) {
		t.Fatalf("expected result to be %v, got %v", expect, result)
	}
}
