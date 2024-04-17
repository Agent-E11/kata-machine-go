package dfsonbst_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/dfsonbst"
	"github.com/agent-e11/kata-machine-go/tests"
)

func TestDFS(t *testing.T) {
	if !dfsonbst.DFS(tests.Tree1, 45) {
		t.Fatal("expected to find 45 in tests.Tree1")
	}
	if !dfsonbst.DFS(tests.Tree1, 7) {
		t.Fatal("expected to find 7 in tests.Tree1")
	}
	if dfsonbst.DFS(tests.Tree1, 69) {
		t.Fatal("expected to find 69 in tests.Tree1")
	}
}
