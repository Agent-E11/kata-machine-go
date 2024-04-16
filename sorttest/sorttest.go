package sorttest

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T, fn func(*[]int)) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	expect := []int{3, 4, 7, 9, 42, 69, 420}

	fn(&arr)

	if !slices.Equal(arr, expect) {
		t.Fatalf("expected slice to be sorted like %v, got %v", expect, arr)
	}

	arr = []int{
		12,  7, 14, 26,  2,
		19, 17,  6, 21,  4,
		16, 10, 30, 27, 23,
		 1,  8,  9,  3, 22,
		25, 11, 29, 13,  5,
		15, 20, 28, 24, 18,
	}
	expect = []int{
		 1,  2,  3,  4,  5,
		 6,  7,  8,  9, 10,
		11, 12, 13, 14, 15,
		16, 17, 18, 19, 20,
		21, 22, 23, 24, 25,
		26, 27, 28, 29, 30,
	}

	fn(&arr)

	if !slices.Equal(arr, expect) {
		t.Fatalf("expected slice to be sorted like %v, got %v", expect, arr)
	}
}
