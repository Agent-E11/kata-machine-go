package mergesort_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/mergesort"
	"github.com/agent-e11/kata-machine-go/sorttest"
)

func TestMergeSort(t *testing.T) {
	sorttest.TestSort(t, mergesort.MergeSort[int])
}
