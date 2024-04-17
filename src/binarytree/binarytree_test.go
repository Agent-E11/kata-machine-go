package binarytree_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/src/binarytree"
	"github.com/agent-e11/kata-machine-go/tests"
)

func TestBFS(t *testing.T) {

	if !binarytree.BFS(tests.Tree1, 45) {
		t.Fatal("expected to find 45 in tests.Tree1")
	}
	if !binarytree.BFS(tests.Tree1, 7) {
		t.Fatal("expected to find 7 in tests.Tree1")
	}
	if binarytree.BFS(tests.Tree1, 69) {
		t.Fatal("expected to not find 69 in tests.Tree1")
	}
}

func TestInOrderSearch(t *testing.T) {
	traversed := binarytree.InOrderTraverse(tests.Tree1)
	expect := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected InOrderSearch(tests.Tree1) to be %v, got %v", expect, traversed)
	}
}

func TestPostOrderSearch(t *testing.T) {
	traversed := binarytree.PostOrderTraverse(tests.Tree1)
	expect := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected PostOrderSearch(tests.Tree1) to be %v, got %v", expect, traversed)
	}
}

func TestPreOrderSearch(t *testing.T) {
	traversed := binarytree.PreOrderTraverse(tests.Tree1)
	expect := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected PreOrderSearch(tests.Tree1) to be %v, got %v", expect, traversed)
	}
}

func TestCompare(t *testing.T) {
	if !binarytree.Compare(&tests.Tree1, &tests.Tree1) {
		t.Fatalf("expected tests.Tree1 to equal tests.Tree1")
	}

	if binarytree.Compare(&tests.Tree1, &tests.Tree2) {
		t.Fatalf("expected tests.Tree1 to not equal tests.Tree2")
	}
}
