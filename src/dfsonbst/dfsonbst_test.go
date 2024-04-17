package dfsonbst_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/dfsonbst"
	"github.com/agent-e11/kata-machine-go/tree"
)

func TestDFS(t *testing.T) {
	if !dfsonbst.DFS(tree.Tree1, 45) {
		t.Fatal("expected to find 45 in tree.Tree1")
	}
	if !dfsonbst.DFS(tree.Tree1, 7) {
		t.Fatal("expected to find 7 in tree.Tree1")
	}
	if dfsonbst.DFS(tree.Tree1, 69) {
		t.Fatal("expected to find 69 in tree.Tree1")
	}
}
