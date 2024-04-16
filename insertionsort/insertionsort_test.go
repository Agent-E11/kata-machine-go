package insertionsort_test

import (
	"testing"

	"github.com/agent-e11/kata-machine-go/insertionsort"
	"github.com/agent-e11/kata-machine-go/sorttest"
)

func TestInsertionSort(t *testing.T) {
	sorttest.TestSort(t, insertionsort.Sort[int])
}
