package prim_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go"
	"github.com/agent-e11/kata-machine-go/tests"
	"github.com/agent-e11/kata-machine-go/src/prim"
)

func TestPrimsList(t *testing.T) {
	result := prim.PrimList(tests.GraphList1)
	expect := kata.WeightedAdjacencyList{
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
		slices.Equal[[]kata.GraphEdge, kata.GraphEdge],
	) {
		t.Fatalf("expected result to be %v, got %v", expect, result)
	}
}
