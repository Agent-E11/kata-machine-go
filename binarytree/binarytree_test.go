package binarytree_test

import (
	"slices"
	"testing"

	bt "github.com/agent-e11/kata-machine-go/binarytree"
	"github.com/agent-e11/kata-machine-go/tree"
)

func TestBFS(t *testing.T) {
	if !bt.BFS(tree.Tree1, 45) {
		t.Fatal("expected to find 45 in tree.Tree1")
	}
	if !bt.BFS(tree.Tree1, 7) {
		t.Fatal("expected to find 7 in tree.Tree1")
	}
	if bt.BFS(tree.Tree1, 69) {
		t.Fatal("expected to not find 69 in tree.Tree1")
	}
}

func TestInOrderSearch(t *testing.T) {
	traversed := bt.InOrderSearch(tree.Tree1)
	expect := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected InOrderSearch(tree.Tree1) to be %v, got %v", expect, traversed)
	}
}

func TestPostOrderSearch(t *testing.T) {
	traversed := bt.PostOrderSearch(tree.Tree1)
	expect := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected PostOrderSearch(tree.Tree1) to be %v, got %v", expect, traversed)
	}
}

func TestPreOrderSearch(t *testing.T) {
	traversed := bt.PreOrderSearch(tree.Tree1)
	expect := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	if !slices.Equal(traversed, expect) {
		t.Fatalf("expected PreOrderSearch(tree.Tree1) to be %v, got %v", expect, traversed)
	}
}
