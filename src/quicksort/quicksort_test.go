package quicksort_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/quicksort"
	"github.com/agent-e11/kata-machine-go/tests/sorttest"
)

func TestQuickSort(t *testing.T) {
	sorttest.TestSort(t, quicksort.Sort[int])
}
