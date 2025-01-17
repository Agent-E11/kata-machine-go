package dfsgraph_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/src/dfsgraph"
	"github.com/agent-e11/kata-machine-go/tests"
)

func TestDFS(t *testing.T) {
	path := dfsgraph.DFSList(tests.GraphList2, 0, 6)
	expect := []int{0, 1, 4, 5, 6}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}

	path = dfsgraph.DFSList(tests.GraphList2, 6, 0)
	expect = []int{}

	if !slices.Equal(path, expect) {
		t.Fatalf("expected path to be %v, got %v", expect, path)
	}
}
