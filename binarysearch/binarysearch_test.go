package binarysearch_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/binarysearch"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	if !binarysearch.Search(list, 69) {
		t.Fatal("expected to find 69 in list")
	}
	if !binarysearch.Search(list, 69420) {
		t.Fatal("expected to find 69420 in list")
	}
	if !binarysearch.Search(list, 1) {
		t.Fatal("expected to find 1 in list")
	}
	if binarysearch.Search(list, 1336) {
		t.Fatal("expected to not find 1336 in list")
	}
	if binarysearch.Search(list, 69421) {
		t.Fatal("expected to not find 69421 in list")
	}
	if binarysearch.Search(list, 0) {
		t.Fatal("expected to not find 0 in list")
	}
}
