package quicksort_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/quicksort"
	"github.com/agent-e11/kata-machine-go/sorttest"
)

func TestQuickSort(t *testing.T) {
	sorttest.TestSort(t, quicksort.Sort[int])
}
