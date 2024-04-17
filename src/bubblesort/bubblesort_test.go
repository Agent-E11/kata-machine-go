package bubblesort_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/src/bubblesort"
	"github.com/agent-e11/kata-machine-go/sorttest"
)

func TestBubbleSort(t *testing.T) {
	sorttest.TestSort(t, bubblesort.Sort[int])
}
