package linearsearchlist_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/linearsearchlist"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	found := linearsearchlist.LinearSearch(arr, 69)
	if !found {
		t.Fatal("expected to find 69 in array")
	}
    found = linearsearchlist.LinearSearch(arr, 1336)
	if found {
		t.Fatal("expected to not find 1336 in array")
	}
    found = linearsearchlist.LinearSearch(arr, 69420)
	if !found {
		t.Fatal("expected to find 69420 in array")
	}
    found = linearsearchlist.LinearSearch(arr, 69421)
	if found {
		t.Fatal("expected to not find 69421 in array")
	}
    found = linearsearchlist.LinearSearch(arr, 1)
	if !found {
		t.Fatal("expected to find 1 in array")
	}
    found = linearsearchlist.LinearSearch(arr, 0)
	if found {
		t.Fatal("expected to not find 0 in array")
	}
}
