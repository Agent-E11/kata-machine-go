package bubblesort_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/kata-machine-go/bubblesort"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	expect := []int{3, 4, 7, 9, 42, 69, 420}

	bubblesort.BubbleSort(&arr)

	if !slices.Equal(arr, expect) {
		t.Fatalf("expected slice to be sorted like %v, got %v", expect, arr)
	}
}
